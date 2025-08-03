import React from 'react';
import type { Expense } from '../types/expense';
import { Card, CardContent, Typography, List, ListItem, IconButton } from '@mui/material';
import DeleteIcon from '@mui/icons-material/Delete';

interface ExpenseListProps {
  expenses: Expense[];
  onSelect?: (expense: Expense) => void;
}

const ExpenseList: React.FC<ExpenseListProps> = ({ expenses, onSelect }) => {
  if (expenses.length === 0) {
    return <Typography>No expenses found.</Typography>;
  }
  return (
    <List>
      {expenses.map(expense => (
        <ListItem key={expense.id} secondaryAction={onSelect ? (
          <IconButton edge="end" aria-label="delete" onClick={() => onSelect(expense)}>
            <DeleteIcon />
          </IconButton>
        ) : null}>
          <Card sx={{ width: '100%' }}>
            <CardContent>
              <Typography variant="h6">{expense.description}</Typography>
              <Typography color="text.secondary">{expense.amount} ({expense.expenseType})</Typography>
              <Typography variant="body2">{new Date(expense.date).toLocaleDateString()}</Typography>
            </CardContent>
          </Card>
        </ListItem>
      ))}
    </List>
  );
};

export default ExpenseList;
