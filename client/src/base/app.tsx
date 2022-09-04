import React from 'react';
import { createRoot } from 'react-dom/client';
import Styles from './index.module.scss';

function App() {
  const [count, setCount] = React.useState(0);

  const add = () => {
    setCount(count + 3);
  };

  return (
    <div className={Styles.base}>
      {count}
      <button onClick={add} type="button">
        +1
      </button>
    </div>
  );
}

createRoot(document.getElementById('app')!).render(<App />);
