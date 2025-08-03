// API service for expenses
import type { Expense, ExpenseInput, ExpenseSearchQuery } from '../types/expense';

const API_BASE = '/api/expenses';

export async function getExpenses(): Promise<Expense[]> {
  const res = await fetch(API_BASE);
  if (!res.ok) throw new Error('Failed to fetch expenses');
  return res.json();
}

export async function getExpenseById(id: number): Promise<Expense> {
  const res = await fetch(`${API_BASE}/${id}`);
  if (!res.ok) throw new Error('Failed to fetch expense');
  return res.json();
}

export async function createExpense(data: ExpenseInput): Promise<Expense> {
  const res = await fetch(API_BASE, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(data),
  });
  if (!res.ok) throw new Error('Failed to create expense');
  return res.json();
}

export async function updateExpense(id: number, data: ExpenseInput): Promise<Expense> {
  const res = await fetch(`${API_BASE}/${id}`, {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(data),
  });
  if (!res.ok) throw new Error('Failed to update expense');
  return res.json();
}

export async function deleteExpense(id: number): Promise<void> {
  const res = await fetch(`${API_BASE}/${id}`, { method: 'DELETE' });
  if (!res.ok) throw new Error('Failed to delete expense');
}

export async function searchExpenses(query: ExpenseSearchQuery): Promise<Expense[]> {
  const res = await fetch(`${API_BASE}/search`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(query),
  });
  if (!res.ok) throw new Error('Failed to search expenses');
  return res.json();
}
