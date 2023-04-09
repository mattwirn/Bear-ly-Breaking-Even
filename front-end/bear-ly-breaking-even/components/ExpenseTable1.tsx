import { useState } from 'react';

function ExpenseTable1() {
  const [rows, setRows] = useState([
    ['Row 1, Column 1', 'Row 1, Column 2'],
    ['Row 2, Column 1', 'Row 2, Column 2'],
    ['Row 3, Column 1', 'Row 3, Column 2'],
  ]);
  
  const addRow = () => {
    const newRow = ['New Row, Column 1', 'New Row, Column 2'];
    setRows([...rows, newRow]);
  };

  return (
    <div className=''>
      <table className='mx-10 px-3 border border-slate-700'>
        <thead>
          <tr>
            <th className='mx-10 px-3 border border-slate-700'>Expense Name</th>
            <th className='mx-10 px-3 border border-slate-700'>Amount</th>
          </tr>
        </thead>
        <tbody>
          {rows.map((row, index) => (
            <tr key={index}>
              <td className='mx-10 px-3  border border-slate-700'>{row[0]}</td>
              <td className='mx-10 px-3 border border-slate-700'>{row[1]}</td>
            </tr>
          ))}
        </tbody>
      </table>
      <button className='mx-10 underline' onClick={addRow}>Add Row</button>
    </div>
  );
}

export default ExpenseTable1;