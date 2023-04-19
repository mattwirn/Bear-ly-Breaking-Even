import { useState, useEffect } from 'react';
import { TrashIcon, PencilIcon } from '@heroicons/react/solid';

const LOCAL_STORAGE_KEY = 'table-rows2';

export default function Table() {
  const [rows, setRows] = useState(() => {
    const storedRows = localStorage.getItem(LOCAL_STORAGE_KEY);
    if (storedRows) {
      return JSON.parse(storedRows);
    }
    return [['', '']];
  });

  useEffect(() => {
    localStorage.setItem(LOCAL_STORAGE_KEY, JSON.stringify(rows));
  }, [rows]);

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

  const getTotal = () => {
    let total = 0;
    rows.forEach((row) => {
      total += parseFloat(row[1]) || 0;
    });
    return total;
  };

  return (
    <div className='mx-10 my-4'>
      <table className=''>
        <thead>
          <td>
            <th className=' px-12'> Expense Name </th>
            <th className=' px-12'> Expense Amount</th>
          </td>
        </thead>
      
        <tbody className='content-center border border-slate-700'>
          {rows.map((row, index) => (
            <tr key={index}>
              <td className='content-center border border-slate-700'>
                <div className='flex justify-between'>
                  <input
                    className='shadow appearance-none border rounded mx-2 my-2 py-1 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline'
                    type='text'
                    value={row[0]}
                    onChange={(e) => editRow(index, 0, e.target.value)}
                  />
                  <input
                    className='shadow appearance-none border rounded mx-2 my-2 py-1 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline'
                    type='number'
                    value={row[1]}
                    onChange={(e) => editRow(index, 1, e.target.value)}
                  />
                  <div className='flex justify-end'>
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
        </tbody>
      </table>
      <table className='grid place-items-left'>
        <tr>
              <button className='mx-1 my-2 underline' onClick={addRow}> Add Row </button>
        </tr>
        <tr className='font-bold '>Total: ${getTotal().toLocaleString('en', {maximumFractionDigits:2 , minimumFractionDigits: 2})}</tr>
      </table>
    </div>
  );
}
