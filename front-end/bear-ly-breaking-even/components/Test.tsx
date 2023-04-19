import Head from 'next/head'
import {useRouter} from "next/router"
import {useState, useEffect} from "react"
import PageHeader from '@/components/PageHeader'
import ExpenseTable from '@/components/ExpenseTable'
import ExpenseTable2 from '@/components/ExpenseTable2'
import ExpenseTable3 from '@/components/ExpenseTable3'
import ExpenseTable4 from '@/components/ExpenseTable4'
import ExpenseTable5 from '@/components/ExpenseTable5'

import classNames from "classnames";


function Tab({ label, active, onClick , href}) {
    return (
      <a
        className={classNames(
          "cursor-pointer",
          "text-gray-600",
          "py-4",
          "px-6",
          "border-b-4",
          active ? "border-blue-500" : ""
        )}
        onClick={onClick}
        href={href}
      >
        {label}
      </a>
    );
  }

  function Test() {
    const [activeTab, setActiveTab] = useState(1);
  
    const renderTable = () => {
      switch (activeTab) {
        case 1:
          return <ExpenseTable />;
        case 2:
          return <ExpenseTable2 />;
        case 3:
          return <ExpenseTable3 />;
        case 4:
          return <ExpenseTable4 />;
        case 5:
          return <ExpenseTable5 />;
        default:
          return null;
      }
    }

  return (
    <div className="justify-center mx-auto px-4">
      <ul className="flex border-b">
        <Tab
          label="Home & Utils"
          href="#1"
          active={activeTab === 1}
          onClick={() => setActiveTab(1)}
        />
        <Tab
          label="Transportation"
          href="#2"
          active={activeTab === 2}
          onClick={() => setActiveTab(2)}
        />
        <Tab
          label="Food & Groceries"
          href="#3"
          active={activeTab === 3}
          onClick={() => setActiveTab(3)}
        />
        <Tab
          label="Entertainment"
          href="#4"
          active={activeTab === 4}
          onClick={() => setActiveTab(4)}
        />
        <Tab
          label="Health"
          href="#5"
          active={activeTab === 5}
          onClick={() => setActiveTab(5)}
        />
      </ul>
      {renderTable()}
    </div>
  );
  }

export default Test;