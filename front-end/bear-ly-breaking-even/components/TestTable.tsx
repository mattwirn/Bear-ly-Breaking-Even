import { useState } from 'react';

function Table() {
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
      <table className='mx-10 border border-slate-700'>
        <thead>
          <tr>
            <th className='mx-10 border border-slate-700'>Column 1</th>
            <th className='mx-10 border border-slate-700'>Column 2</th>
          </tr>
        </thead>
        <tbody>
          {rows.map((row, index) => (
            <tr key={index}>
              <td className='mx-10 border border-slate-700'>{row[0]}</td>
              <td className='mx-10 border border-slate-700'>{row[1]}</td>
            </tr>
          ))}
        </tbody>
      </table>
      <button className='mx-10 underline' onClick={addRow}>Add Row</button>
    </div>
  );
}

export default Table;