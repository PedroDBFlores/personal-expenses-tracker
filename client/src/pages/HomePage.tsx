import React from 'react';
import { Container, Typography, Box, Button } from '@mui/material';
import { Link } from 'react-router-dom';

const HomePage: React.FC = () => (
  <Container maxWidth="sm" sx={{ mt: 8, textAlign: 'center' }}>
    <Box>
      <Typography variant="h3" gutterBottom>Personal Expenses Tracker</Typography>
      <Typography variant="h6" gutterBottom>Track your expenses easily and efficiently.</Typography>
      <Button variant="contained" color="primary" component={Link} to="/expenses" sx={{ mt: 4 }}>
        Go to Expenses
      </Button>
    </Box>
  </Container>
);

export default HomePage;
