import React from 'react';
import { BrowserRouter as Router, Route } from 'react-router-dom'

import './App.css';
import Navbar from './components/Navbar';
import Projects from './views/Projects'

function App() {
  return (
    <div className="App">
      <Navbar />
      <Router>
        <Route path="/" exact component={Projects} />
      </Router>
    </div>
  );
}

export default App;
