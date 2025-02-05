
import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Home from './pages/Home';
import Search from './pages/Search';
import Log from './pages/Log';
import Signup from './pages/Signup';
import Cont from './pages/Cont';
import Profile from './pages/Profile';
import Aboutpage from './pages/Aboutpage';
import Cart from './components/Cart';
import ShopPage from './pages/ShopPage';
import CateProduct from './pages/BookPage';
import Pen from './pages/PenPage';
import Paper from './pages/PaperPage';
import DrawingBook from './pages/DrawingBookPage';
import DIY from './pages/DIY_Page';
import Shop1 from './pages/Shop1_Page';
import Shop2 from './pages/Shop2_Page';
import Shop3 from './pages/Shop3_Page';

import './App.css';




function App() {
    return (
        <Router>
            <div className="App">
                <Routes>
                    <Route path="/" element={<Home />} />
                    <Route path="/search" element={<Search />} />
                    <Route path="/login" element={<Log />} />
                    <Route path="/signup" element={<Signup />} />
                    <Route path="/contact" element={<Cont />} />
                    <Route path="/profile" element={<Profile />} /> 
                    <Route path="/Aboutpage" element={<Aboutpage />} />
                    <Route path="/cart" element={<Cart />} />
                    <Route path="/product" element={<ShopPage />} />
                    <Route path="/Book" element={<CateProduct />} />
                    <Route path="/pen" element={<Pen />} />
                    <Route path="/paper" element={<Paper />} />
                    <Route path="/drawingbook" element={<DrawingBook />} />
                    <Route path="/diy" element={<DIY />} />
                    <Route path="/seller" element={<Shop1 />} />
                    <Route path="/seller" element={<Shop2 />} />
                    <Route path="/seller" element={<Shop3 />} />
                    
                    
                    </Routes>
            </div>
        </Router>
    );
}

export default App;     
        
      

        
  
                    
                    

  