import Head from 'next/head'
import {useRouter} from "next/router"
import {useState, useEffect} from "react"
import PageHeader from '@/components/PageHeader'
import ExpenseTable from '@/components/ExpenseTable'
import classNames from "classnames";


function Tab({ label, active, onClick }) {
    return (
      <li
        className={classNames(
          "cursor-pointer",
          "text-gray-600",
          "py-4",
          "px-6",
          "border-b-4",
          active ? "border-blue-500" : ""
        )}
        onClick={onClick}
      >
        {label}
      </li>
    );
  }

  function Test() {
    const [activeTab, setActiveTab] = useState(1);
  
    const renderTable = () => {
      switch (activeTab) {
        case 1:
          return <ExpenseTable />;
        case 2:
          return <ExpenseTable />;
        case 3:
          return <ExpenseTable />;
        case 4:
          return <ExpenseTable />;
        case 5:
          return <ExpenseTable />;
        default:
          return null;
      }
    }

  return (
    <div className="container mx-auto px-4">
      <ul className="flex border-b">
        <Tab
          label="Table 1"
          active={activeTab === 1}
          onClick={() => setActiveTab(1)}
        />
        <Tab
          label="Table 2"
          active={activeTab === 2}
          onClick={() => setActiveTab(2)}
        />
        <Tab
          label="Table 3"
          active={activeTab === 3}
          onClick={() => setActiveTab(3)}
        />
        <Tab
          label="Table 4"
          active={activeTab === 4}
          onClick={() => setActiveTab(4)}
        />
        <Tab
          label="Table 5"
          active={activeTab === 5}
          onClick={() => setActiveTab(5)}
        />
      </ul>
      {renderTable()}
    </div>
  );
  }

export default Test;