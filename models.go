package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/lancatlin/lazy-finance/ledger"
	"github.com/lancatlin/lazy-finance/model"
)

type User struct {
	IsLogin  bool
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func (u *User) Dir() string {
	dir := path.Join(config.DataPath, u.Email)
	return dir
}

func (u *User) Mkdir() error {
	err := os.Mkdir(u.Dir(), 0755)
	if err != nil {
		return fmt.Errorf("failed to create user directory: %w", err)
	}
	files, err := os.ReadDir(ARCHETYPES_DIR)
	if err != nil {
		return fmt.Errorf("failed to read archetypes directory: %w", err)
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		err := u.copyArchetypeFile(file.Name())
		if err != nil {
			return fmt.Errorf("failed to copy archetype file: %w", err)
		}
	}
	return nil
}

func (u *User) copyArchetypeFile(file string) error {
	src, err := os.OpenFile(path.Join(ARCHETYPES_DIR, file), os.O_RDONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open archetype file: %w", err)
	}
	defer src.Close()

	dest, err := u.WriteFile(file)
	if err != nil {
		return fmt.Errorf("failed to open destination file: %w", err)
	}

	_, err = io.Copy(dest, src)
	if err != nil {
		return fmt.Errorf("failed to copy file: %w", err)
	}
	return dest.Close()
}

func (u *User) FilePath(name string) string {
	return path.Join(u.Dir(), name)
}

func (u *User) File(name string, mode int) (*os.File, error) {
	return os.OpenFile(u.FilePath(name), mode, 0644)
}

func (u *User) AppendFile(name string) (*os.File, error) {
	return u.File(name, os.O_WRONLY|os.O_CREATE|os.O_APPEND)
}

func (u *User) ReadFile(name string) (*os.File, error) {
	return u.File(name, os.O_RDONLY|os.O_CREATE)
}

func (u *User) WriteFile(name string) (*os.File, error) {
	return u.File(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC)
}

func (u *User) ListFiles() ([]File, error) {
	files, err := os.ReadDir(u.Dir())
	if err != nil {
		return nil, fmt.Errorf("failed to open directory: %w", err)
	}
	result := make([]File, len(files))
	for i, v := range files {
		result[i] = File{
			Name: v.Name(),
		}
	}
	return result, nil
}

func (u *User) readAllFile(name string) (data []byte, err error) {
	f, err := u.ReadFile(name)
	if err != nil {
		return
	}
	defer f.Close()
	data, err = io.ReadAll(f)
	return
}

func (u *User) appendToFile(tx string) (err error) {
	f, err := u.AppendFile(DEFAULT_JOURNAL)
	if err != nil {
		return
	}

	buf := strings.NewReader(strings.ReplaceAll(tx, "\r", "")) // Remove CR generated from browser
	_, err = io.Copy(f, buf)
	if err != nil {
		return
	}
	return f.Close()
}

func (u *User) overwriteFile(filename string, tx string) (err error) {
	f, err := u.WriteFile(filename)
	if err != nil {
		return err
	}

	buf := strings.NewReader(strings.ReplaceAll(tx, "\r", "")) // Remove CR generated from browser
	_, err = io.Copy(f, buf)
	if err != nil {
		return err
	}
	return f.Close()
}

func (u *User) templates() (templates []model.Template, err error) {
	f, err := u.ReadFile("templates.json")
	if err != nil {
		return nil, fmt.Errorf("failed to read templates.json: %w", err)
	}
	defer f.Close()

	var buf bytes.Buffer
	_, err = buf.ReadFrom(f)
	if err != nil {
		return nil, fmt.Errorf("failed to read from buffer: %w", err)
	}

	templates, err = model.LoadTemplates(buf.String())
	if err != nil {
		return nil, fmt.Errorf("failed to load templates: %w", err)
	}
	return templates, nil
}

func (u *User) newTx(tx model.Transaction) error {
	if err := tx.Validate(); err != nil {
		return err
	}
	txString, err := tx.Generate()
	if err != nil {
		return err
	}
	err = u.appendToFile(txString)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) txs(query ledger.Query) ([]model.Transaction, error) {
	f, err := u.ReadFile(DEFAULT_JOURNAL)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	query.Command = "reg"
	command := ledger.NewCommand(query, u.Dir(), f)
	output, err := command.Execute()
	if err != nil {
		return nil, err
	}
	return ledger.LoadTransactions(output)
}

func (u *User) getBalances(query ledger.Query) (balances []ledger.Balance, err error) {
	f, err := u.ReadFile(DEFAULT_JOURNAL)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	query.Command = "bal"
	command := ledger.NewCommand(query, u.Dir(), f)
	output, err := command.Execute()
	if err != nil {
		return nil, err
	}
	return ledger.LoadBalances(output)
}

func (u *User) saveTemplate(template model.Template) error {
	templates, err := u.templates()
	if err != nil {
		return err
	}
	templates = append(templates, template)

	return u.saveTemplates(templates)
}

func (u *User) saveTemplates(templates []model.Template) error {
	data, err := json.Marshal(templates)
	if err != nil {
		return err
	}
	return u.overwriteFile("templates.json", string(data))
}
