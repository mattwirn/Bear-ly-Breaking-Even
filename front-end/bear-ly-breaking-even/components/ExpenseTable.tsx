import { useState } from 'react';
import { TrashIcon, PencilIcon } from '@heroicons/react/solid';

export default function Table() {
  const [rows, setRows] = useState([
    ['Row 1, Column 1 table 1', 'Row 1, Column 2']

  ]);

  const addRow = () => {
    const newRow = ['', ''];
    setRows([...rows, newRow]);
  };

  const deleteRow = (index) => {
    const newRows = [...rows];
    newRows.splice(index, 1);
    setRows(newRows);
  };

  const editRow = (index, colIndex, value) => {
    const newRows = [...rows];
    newRows[index][colIndex] = value;
    setRows(newRows);
  };

  function showInputFields() {

  }

  return (
    <div className='mx-10 my-4'>
      <table className=' border border-black-700'>
        <tbody>
          {rows.map((row, index) => (
            <tr key={index}>
              <td className='content-center border border-slate-700'>
                <th className= 'px-10 '> Expense Name </th>
                <th className= 'px-10 '> Expense Amount</th>
                    <div className='flex justify-between'>
                  <input
                    className='mx-3 my-2 border border-slate-700'
                    type='text'
                    value={row[0]}
                    onChange={(e) => editRow(index, 0, e.target.value)}
                  />
                  <input
                    className='my-2 border border-slate-700'
                    type='text'
                    value={row[1]}
                    onChange={(e) => editRow(index, 1, e.target.value)}
                  />
                  <div className='flex justify-end'>
                  <button onClick={showInputFields} className='my-2'>
                      <PencilIcon className='h-4 w-4 text-blue-500' />
                    </button>
                    <button
                      className=' my-2'
                      onClick={() => deleteRow(index)}
                    >
                      <TrashIcon className='h-4 w-4 text-red-500' />
                    </button>
                  </div>
                </div>
              </td>
            </tr>
          ))}
          <tr>
            <td className='border border-black-700'>
              <button className='mx-1 my-2 hover:underline' onClick={addRow}>
                Add Row
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  );
}
