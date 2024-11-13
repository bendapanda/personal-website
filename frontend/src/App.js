import Navbar from './components/Navbar';
import {
  BrowserRouter as Router,
  Routes,
  Route, 
} from 'react-router-dom'

import Home from "./pages/Home"
import Resume from './pages/Resume';

function App() {
  return (
    <Router>
      <Navbar />
      <Routes>
        <Route exact path='/' element= {<Home />} />
        <Route exact path="/resume" element= {<Resume/>}/>
      </Routes>
    </Router>
  );
}

export default App;
