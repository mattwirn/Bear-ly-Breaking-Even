import { useState, useEffect } from 'react';
import { TrashIcon, PencilIcon } from '@heroicons/react/solid';

const LOCAL_STORAGE_KEY = 'table-rows1';

export default function Table() {
  const [rows, setRows] = useState([['', '']]);

  useEffect(() => {
    const storedRows = localStorage.getItem(LOCAL_STORAGE_KEY);

    if (storedRows) {
      setRows(JSON.parse(storedRows));
    }
  }, []);

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
      total += parseInt(row[1]) || 0;
    });
    return total;
  };

  return (
    <div className='mx-auto my-4'>
      <table className=' w-1/2 border border-black'>
        <tbody>
          <tr className='mx-auto'>
            <th className=''> Expense Name </th>
            <th className=''> Expense Amount</th>
          </tr>
          {rows.map((row, index) => (
            <tr key={index}>
              <td className='flex justify-between border border-black'>
                <input
                  className='mx-3 my-2'
                  type='text'
                  value={row[0]}
                  onChange={(e) => editRow(index, 0, e.target.value)}
                />
                <input
                  className='my-2'
                  type='text'
                  value={row[1]}
                  onChange={(e) => editRow(index, 1, e.target.value)}
                />
                <div className='flex justify-end'>
                  <button className='my-2'>
                    <PencilIcon className='h-4 w-4 text-blue-500' />
                  </button>
                  <button
                    className=' my-2'
                    onClick={() => deleteRow(index)}
                  >
                    <TrashIcon className='h-4 w-4 text-red-500' />
                  </button>
                </div>
              </td>
            </tr>
          ))}
          <tr>
            <td className=''>
              <button className='mx-1 my-2 hover:underline' onClick={addRow}>
                Add Row
              </button>
            </td>
          </tr>
        </tbody>
      </table>
      Total: ${getTotal()}
    </div>
  );
}
