import React from 'react'

var thing1 = "mortgage and shi"
var thing2 = 69
var n = 1
export default function Table1() {
    return (
        <table className=" mx-10 border-collapse border border-slate-500 ...">
        <thead>
          <tr>
            <th className="border border-slate-600">Expense Name</th>
            <th className="border border-slate-600 ...">Amount</th>
          </tr>
        </thead>
        <tbody id="container">
          <tr>
            <td className="mx-10 border border-slate-700 ...">{thing1}</td>
            <td className="border border-slate-700 ...">{thing2}</td>
          </tr>
          {read()}
          <tr>
            <td className="border border-slate-700 ...">Ohio</td>
            <td className="border border-slate-700 ...">Columbus</td>
          </tr>

          <tr>
            <td className="border border-slate-700 ...">Michigan</td>
            <td className="border border-slate-700 ...">Detroit</td>
          </tr>
        </tbody>
      </table>
    )
}
function read() {
    var result= ''
    for (let i = 0; i < 6; i++) {
        return (
            <tr>
            <td className="mx-10 border border-slate-700 ...">{thing1}</td>
            <td className="border border-slate-700 ...">{thing2}</td>
            </tr>
        )
    }
    return (result)
    
}
function printRow(){
    return (
        <tr>
        <td className="mx-10 border border-slate-700 ...">{thing1}</td>
        <td className="border border-slate-700 ...">{thing2}</td>
        </tr>
    )
}
