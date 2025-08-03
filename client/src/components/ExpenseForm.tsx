import React, { useState } from 'react';
import type { ExpenseInput, Expense, ExpenseType } from '../types/expense';
import { TextField, Button, MenuItem, Box } from '@mui/material';

interface ExpenseFormProps {
  initial?: ExpenseInput | Expense;
  onSubmit: (data: ExpenseInput) => void;
  submitLabel?: string;
}

const defaultExpense: ExpenseInput = {
  amount: 0,
  description: '',
  date: new Date().toISOString().slice(0, 10),
  expenseType: 'debit',
  fulfillsExpenseId: undefined,
};

const ExpenseForm: React.FC<ExpenseFormProps> = ({ initial, onSubmit, submitLabel = 'Save' }) => {
  const [form, setForm] = useState<ExpenseInput>({ ...defaultExpense, ...initial });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setForm(f => ({ ...f, [name]: name === 'amount' ? Number(value) : value }));
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    onSubmit(form);
  };

  return (
    <Box component="form" onSubmit={handleSubmit} sx={{ display: 'flex', flexDirection: 'column', gap: 2 }}>
      <TextField
        label="Description"
        name="description"
        value={form.description}
        onChange={handleChange}
        required
      />
      <TextField
        label="Amount"
        name="amount"
        type="number"
        value={form.amount}
        onChange={handleChange}
        required
      />
      <TextField
        label="Date"
        name="date"
        type="date"
        value={form.date}
        onChange={handleChange}
        InputLabelProps={{ shrink: true }}
        required
      />
      <TextField
        select
        label="Type"
        name="expenseType"
        value={form.expenseType}
        onChange={handleChange}
        required
      >
        <MenuItem value="debit">Debit</MenuItem>
        <MenuItem value="credit">Credit</MenuItem>
      </TextField>
      <TextField
        label="Fulfills Expense ID"
        name="fulfillsExpenseId"
        type="number"
        value={form.fulfillsExpenseId ?? ''}
        onChange={handleChange}
        InputLabelProps={{ shrink: true }}
      />
      <Button type="submit" variant="contained">{submitLabel}</Button>
    </Box>
  );
};

export default ExpenseForm;
