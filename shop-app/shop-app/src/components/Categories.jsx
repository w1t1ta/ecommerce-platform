

import React from 'react';
import { Link } from 'react-router-dom';
import './css/Categories.css';

const categories = [
    { name: "Book", imageUrl: "https://img.freepik.com/free-vector/hand-drawn-flat-design-stack-books-illustration_23-2149341898.jpg", link: "/book" },
    { name: "Pen/Pencil", imageUrl: "https://i5.walmartimages.com/seo/Paper-Mate-Flair-Felt-Tip-Pens-Medium-Tip-Limited-Edition-24-Count_6b6d80f8-38c2-4241-98a2-5639cc55bd5e.cc8a5c61fabb5cde40df81ea8824e4a6.jpeg", link: "/pen" },
    { name: "Paper", imageUrl: "https://www.sswholesale.com.au/website/var/tmp/image-thumbnails/20000/20927/thumb__sswProductImagesFancybox/Cover-Papers.jpeg", link: "/paper" },
    { name: "Drawing Book", imageUrl: "https://m.media-amazon.com/images/I/71ueySfoI6L.jpg", link: "/drawingbook" },
    { name: "DIY", imageUrl: "https://images.squarespace-cdn.com/content/v1/63dde481bbabc6724d988548/240ad38c-edfa-4611-85e1-6bc1e27230a9/_116d2caf-fdfb-4e8a-af69-02d65dee6bc7.jpg", link: "/diy" },
];


const Category = () => {
    return (
        <div className="category-container">

            {/* Browse by Category */}
            <h2 className="category-heading">Browse by Category</h2>
            <div className="category">
                {categories.map((category, index) => (
                    <div key={index} className="category-item">
                        <Link to={category.link}>
                            <img src={category.imageUrl} alt={category.name} />
                            <div className="category-name">{category.name}</div>
                        </Link>
                    </div>
                ))}
            </div>
        </div>
    );
};

export default Category;
