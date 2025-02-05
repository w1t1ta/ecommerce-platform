import React from 'react';
import Header from '../components/Header';
import Banner from '../components/Banner';
import Categories from '../components/Categories';
import Footer from '../components/Footer';
import NewProducts from '../components/Newproduct';
import SellerList from '../components/SellerList';
import Recommend from '../components/Recommend';




const Home = () => {
    return (
        <div>
            <Header />
            <Banner />
            <NewProducts/>
            <Recommend/>
            <SellerList/>
            <Categories />
            <Footer />
        </div>
    );
}

export default Home;