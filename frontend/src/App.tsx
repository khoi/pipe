import { useState } from "react";
import logo from "./assets/images/logo-universal.png";
import "./App.css";

function App() {
  const [resultText, setResultText] = useState(
    "Please enter your name below 👇"
  );
  const [name, setName] = useState("");
  const updateName = (e: any) => setName(e.target.value);
  const updateResultText = (result: string) => setResultText(result);

  return <div id="App">Hello</div>;
}

export default App;
