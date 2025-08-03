import React, { useEffect, useState } from 'react';
import { getExpenses, createExpense, deleteExpense } from '../services/expenseService';
import type { Expense, ExpenseInput } from '../types/expense';
import ExpenseList from '../components/ExpenseList';
import ExpenseForm from '../components/ExpenseForm';
import { Container, Typography, Snackbar, Alert, Paper } from '@mui/material';

const ExpensesPage: React.FC = () => {
  const [expenses, setExpenses] = useState<Expense[]>([]);
  const [error, setError] = useState<string | null>(null);
  const [success, setSuccess] = useState<string | null>(null);

  const fetchExpenses = async () => {
    try {
      setExpenses(await getExpenses());
    } catch (e) {
      setError((e as Error).message);
    }
  };

  useEffect(() => {
    fetchExpenses();
  }, []);

  const handleAdd = async (data: ExpenseInput) => {
    try {
      await createExpense(data);
      setSuccess('Expense added!');
      fetchExpenses();
    } catch (e) {
      setError((e as Error).message);
    }
  };

  const handleDelete = async (expense: Expense) => {
    try {
      await deleteExpense(expense.id);
      setSuccess('Expense deleted!');
      fetchExpenses();
    } catch (e) {
      setError((e as Error).message);
    }
  };

  // For simplicity, update is not shown in UI, but can be added similarly

  return (
    <Container maxWidth="md" sx={{ mt: 4 }}>
      <Typography variant="h4" gutterBottom>Expenses</Typography>
      <Paper sx={{ p: 2, mb: 4 }}>
        <ExpenseForm onSubmit={handleAdd} submitLabel="Add Expense" />
      </Paper>
      <ExpenseList expenses={expenses} onSelect={handleDelete} />
      <Snackbar open={!!error} autoHideDuration={4000} onClose={() => setError(null)}>
        <Alert severity="error">{error}</Alert>
      </Snackbar>
      <Snackbar open={!!success} autoHideDuration={2000} onClose={() => setSuccess(null)}>
        <Alert severity="success">{success}</Alert>
      </Snackbar>
    </Container>
  );
};

export default ExpensesPage;
