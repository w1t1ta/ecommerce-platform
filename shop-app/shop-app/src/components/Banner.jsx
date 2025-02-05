import React, { useState, useEffect } from 'react';
import { Carousel } from 'react-bootstrap';
import './css/Banner.css';

const Banner = () => {
  const [currentIndex, setCurrentIndex] = useState(0);
  const images = [
    "https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEguusqjegfCxSA3GZcTRy6iL_wWKD9h-vSmrJQujwBCHziVZ82_UeaARPK7I53py09LXP423wQU2719NXlYCsc4Bb1fUHneghFcGsEPqsXvzI1F1YIe77-7RERfaQ_sBi6p_5_8TQ-Ud12wjBwTubXAz4yvno5W4B49kxyN5pbSp4N1ebSmPaqQmEOuljM/s16000/1.png",
    "https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEi6Q417B6MYpDRUEqHQWBe-B-FTGYFNzr1my29QZzl-B73tdvWLgvSEmtdjzwmTlnt_QWIbPzLyPiPb0jQUFqtcsyjRdmBsUkCMqsU0_57dMFyMNbRl7MKlUNNXn_nFWLSOcJx-CuOKRRtSIOnLu6ZaTzDQTypxKF5F5L7kVzPyT-SY6tADmKa7dfq0044/s16000/2.png",
    "https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEiJ_PCDVFrIJ_ubYvSm7MFa8QTR1P9ImAYwGVICsI1uNbihEQM79jQc3WGL3ncoms3IBWgrTspdZY3hDm8BZYXoMHdDwJczbEmuFj6LvU9Mcf0Ndi4cT83HnsGCVIW19ZBQD7pGpHQdIobjgcnDEMbQlQjS3Aj-mNO8BPl69XySaa5AEdkbHuEgY88pon0/s16000/3.png"
  ];

  // ฟังก์ชันสำหรับเลื่อนไปยังสไลด์ถัดไป
  const nextSlide = () => {
    setCurrentIndex((prevIndex) => (prevIndex + 1) % images.length);
  };

  // ฟังก์ชันสำหรับเลื่อนไปยังสไลด์ก่อนหน้า
  const prevSlide = () => {
    setCurrentIndex((prevIndex) => (prevIndex - 1 + images.length) % images.length);
  };

  // ตั้งเวลาให้เลื่อนอัตโนมัติทุก ๆ 3 วินาที
  useEffect(() => {
    const interval = setInterval(nextSlide, 3000);
    return () => clearInterval(interval);
  }, []);

  return (
    <div>
      <Carousel activeIndex={currentIndex} onSelect={(index) => setCurrentIndex(index)}>
        {images.map((image, index) => (
          <Carousel.Item key={index}>
            <img className="d-block w-100" src={image} alt={`Slide ${index + 1}`} />
          </Carousel.Item>
        ))}
      </Carousel>
      <button className="prev" onClick={prevSlide}>&#10094;</button>
      <button className="next" onClick={nextSlide}>&#10095;</button>
    </div>
  );
};

export default Banner;