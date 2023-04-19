import React, { useState, useEffect } from 'react';

interface Expense {
  name: string;
  amount: number;
}

const ExpenseTracker: React.FC = () => {
  const [expenses, setExpenses] = useState<Expense[]>([{ name: '', amount: 0 }]);
  const [totalAmount, setTotalAmount] = useState<number>(0);

  useEffect(() => {
    const storedExpenses = localStorage.getItem('expenses');
    if (storedExpenses) {
      setExpenses(JSON.parse(storedExpenses));
    }
  }, []);

  useEffect(() => {
    localStorage.setItem('expenses', JSON.stringify(expenses));
  }, [expenses]);

  const addExpenseRow = () => {
    setExpenses([...expenses, { name: '', amount: 0 }]);
  };

  const handleExpenseChange = (index: number, field: keyof Expense, value: string | number) => {
    const updatedExpenses = [...expenses];
    updatedExpenses[index][field] = value;
    setExpenses(updatedExpenses);
  };

  const calculateTotalAmount = () => {
    const amounts = expenses.map(expense => expense.amount);
    const total = amounts.reduce((acc, curr) => acc + curr, 0);
    setTotalAmount(total);
  };

  return (
    <div>
      <h1>Expense Tracker</h1>
      <table>
        <thead>
          <tr>
            <th>Expense Name</th>
            <th>Expense Amount</th>
          </tr>
        </thead>
        <tbody>
          {expenses.map((expense, index) => (
            <tr key={index}>
              <td>
                <input
                  type="text"
                  value={expense.name}
                  onChange={(event) => handleExpenseChange(index, 'name', event.target.value)}
                />
              </td>
              <td>
                <input
                  type="number"
                  value={expense.amount}
                  onChange={(event) => handleExpenseChange(index, 'amount', parseFloat(event.target.value))}
                />
              </td>
            </tr>
          ))}
        </tbody>
      </table>
      <br />
      <button onClick={addExpenseRow}>Add Expense</button>
      <button onClick={calculateTotalAmount}>Calculate Total</button>
      <p>Total Expense Amount: {totalAmount}</p>
    </div>
  );
};

export default ExpenseTracker;
