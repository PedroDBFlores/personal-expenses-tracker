import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';
import HomePage from './pages/HomePage';
import ExpensesPage from './pages/ExpensesPage';

function App() {
  return (
    <Router>
      <nav style={{ padding: '1rem', textAlign: 'center' }}>
        <Link to="/">Home</Link> | <Link to="/expenses">Expenses</Link>
      </nav>
      <Routes>
        <Route path="/" element={<HomePage />} />
        <Route path="/expenses" element={<ExpensesPage />} />
      </Routes>
    </Router>
  );
}

export default App;
