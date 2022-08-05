import './App.css';
import Header from './components/Header';
import Top from './components/Top';
import Footer from './components/Footer';
import Login from './components/Login';
import Logout from './components/Logout';
import { BrowserRouter, Routes, Route} from 'react-router-dom';

function App() {
  return (
    <>
      <BrowserRouter>
        <Header />
        <Routes>
          <Route exact path='/' element={<Top/>} />
          <Route path='/login' element={<Login/>} />
          <Route path='/logout' element={<Logout/>} />
        </Routes>
        <Footer />
      </BrowserRouter>
    </>
  );
}

export default App;
