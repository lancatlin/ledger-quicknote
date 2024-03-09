export interface Account {
  name: string;
  amount: number;
  commodity: string;
}

export interface Template {
  name: string;
  accounts: Account[];
}

export interface Transaction {
  name: string;
  date: Date;
  accounts: Account[];
}

export const defaultTemplate: Template = {
  name: "default",
  accounts: [
    {
      name: "",
      amount: 0,
      commodity: "",
    },
    {
      name: "",
      amount: 0,
      commodity: "",
    },
  ],
};

export interface Balance {
  account: string;
  balance: string;
}
