import React from 'react';
import './css/Footer.css'

const Footer = () => {
    return (
        <footer className="footer">
            <div className="footer-column">
                <div className="column-head">
                    <h4>Support</h4>
                </div>
                <ul>
                    <h5>6 Ratchamakkha Nai Road,<br />
                    Mueang District, Nakhon Pathom 73000<br />
                    Main phone number: +66 34 109 686,<br />
                    press 0 for operator</h5>
                </ul>
            </div>
            <div className="footer-column">
                <div className="column-head">
                    <h4>Quick Link</h4>
                </div>
                <ul>
                    <li><a href="https://club.b2s.co.th/th/privacy-and-policy.php">Policy</a></li> {/* ขออนุญาตลิงค์หน้าตัวอย่าง policy ของ B2S*/}
                    <li><a href="https://club.b2s.co.th/th/faq.php">FAQ</a></li> {/* ขออนุญาตลิงค์หน้าตัวอย่าง FAQ ของ B2S*/}
                    <li><a href="#Contact">Contact</a></li> {/* ลิงค์หน้า contact ของนิว */}
                </ul>
            </div>
            <div className="footer-column">
                <div className="column-head">
                    <h4>About Us</h4>
                </div>
                <ul>
                    <li><a href="https://web.facebook.com/B2SThailand/?locale=th_TH&_rdc=1&_rdr"><img src="https://uxwing.com/wp-content/themes/uxwing/download/brands-and-social-media/facebook-app-round-white-icon.png" 
                    style={{ width: '15px' }} /></a></li> {/* ขออนุญาตลิงค์ facebook ของ B2S*/}
                    <li><a href="https://www.instagram.com/b2sthailand/"><img src="https://img.icons8.com/win10/512/FFFFFF/instagram-new.png" 
                    style={{ width: '20px' }} /></a></li> {/* ขออนุญาตลิงค์ instagram ของ B2S*/}
                    <li><a href="https://x.com/b2sthailand?lang=en"><img src="https://uxwing.com/wp-content/themes/uxwing/download/brands-and-social-media/x-social-media-white-icon.png" 
                    style={{ width: '15px' }} /></a></li> {/* ขออนุญาตลิงค์ X ของ B2S*/}
                </ul>
            </div>
        </footer>
    );
};

export default Footer;