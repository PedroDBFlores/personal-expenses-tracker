// Expense type definitions for the client

export type ExpenseType = 'credit' | 'debit';

export interface Expense {
  id: number;
  amount: number;
  description: string;
  date: string; // ISO string
  createdAt: string;
  updatedAt: string;
  expenseType: ExpenseType;
  fulfillsExpenseId?: number | null;
}

export interface ExpenseInput {
  amount: number;
  description: string;
  date: string;
  expenseType: ExpenseType;
  fulfillsExpenseId?: number | null;
}

export interface ExpenseSearchQuery {
  dateFrom?: string;
  dateTo?: string;
  amountMin?: number;
  amountMax?: number;
  expenseType?: ExpenseType;
}
