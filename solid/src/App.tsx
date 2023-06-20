import type { Component } from "solid-js";
import styles from "./App.module.css";
import { createSignal } from "solid-js";

const App: Component = () => {
  const [keyword, setKeyword] = createSignal("");

  const handleKeywordChange = (e) => {
    setKeyword(e.target.value);
  };

  const submitKeyword = () => {
    console.log(keyword());
  };

  return (
    <div class={styles.App}>
      <header class={styles.header}>
        <p>Scrape for deals</p>

        <input
          type="text"
          onInput={(e) => handleKeywordChange(e)}
          placeholder="Whatcha looking for?"
        />
        <button type="button" onClick={submitKeyword}>
          Search
        </button>
      </header>
    </div>
  );
};

export default App;
