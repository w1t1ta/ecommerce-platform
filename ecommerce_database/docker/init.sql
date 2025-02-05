-- ตรวจสอบว่า database มีอยู่แล้วหรือไม่ และสร้างถ้ายังไม่มี
DO
$$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_database WHERE datname = 'ecommerce') THEN
        PERFORM dblink_exec('dbname=' || current_database(), 'CREATE DATABASE ecommerce');
    END IF;
END
$$;

-- เชื่อมต่อกับ Database ที่สร้าง
\c ecommerce;

-- สร้าง Extension สำหรับ UUID (เฉพาะ PostgreSQL)
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- สร้าง ENUM สำหรับ product_status และ product_type
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'product_status') THEN
        CREATE TYPE product_status AS ENUM ('active', 'inactive');
    END IF;
END
$$;

-- สร้างตาราง categories
CREATE TABLE IF NOT EXISTS categories (
    category_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    parent_category_id INTEGER,
    FOREIGN KEY (parent_category_id) REFERENCES categories(category_id) ON DELETE SET NULL
);

-- สร้างตาราง sellers
CREATE TABLE IF NOT EXISTS sellers (
    seller_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL UNIQUE,
    contact_info VARCHAR(255),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- สร้างตาราง products
CREATE TABLE IF NOT EXISTS products (
    product_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    brand VARCHAR(255),
    model_number VARCHAR(100),
    sku VARCHAR(100) UNIQUE NOT NULL,
    price NUMERIC(10, 2) NOT NULL CHECK (price >= 0),
    status product_status NOT NULL,
    seller_id UUID NOT NULL,
    category_id INTEGER NOT NULL,
    product_type VARCHAR(100) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (seller_id) REFERENCES sellers(seller_id) ON DELETE CASCADE,
    FOREIGN KEY (category_id) REFERENCES categories(category_id) ON DELETE CASCADE
);

-- สร้างตาราง books
CREATE TABLE IF NOT EXISTS books (
    book_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    product_id UUID NOT NULL UNIQUE,
    author VARCHAR(255) NOT NULL,
    publisher VARCHAR(255) NOT NULL,
    isbn VARCHAR(20) UNIQUE NOT NULL,
    publication_date DATE NOT NULL,
    number_of_pages INTEGER NOT NULL CHECK (number_of_pages > 0),
    FOREIGN KEY (product_id) REFERENCES products(product_id) ON DELETE CASCADE
);

-- สร้างตาราง pencil
CREATE TABLE IF NOT EXISTS pencil (
    pencil_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    product_id UUID NOT NULL UNIQUE,
    material VARCHAR(100) NOT NULL,
    FOREIGN KEY (product_id) REFERENCES products(product_id) ON DELETE CASCADE
);

-- สร้างตาราง paper
CREATE TABLE IF NOT EXISTS paper (
    paper_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    product_id UUID NOT NULL UNIQUE,
    material VARCHAR(100) NOT NULL,
    FOREIGN KEY (product_id) REFERENCES products(product_id) ON DELETE CASCADE
);

-- สร้างตาราง notebook
CREATE TABLE IF NOT EXISTS notebook (
    notebook_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    product_id UUID NOT NULL UNIQUE,
    material VARCHAR(100) NOT NULL,
    FOREIGN KEY (product_id) REFERENCES products(product_id) ON DELETE CASCADE
);

-- สร้างตาราง DIY
CREATE TABLE IF NOT EXISTS DIY (
    DIY_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    product_id UUID NOT NULL UNIQUE,
    material VARCHAR(100) NOT NULL,
    FOREIGN KEY (product_id) REFERENCES products(product_id) ON DELETE CASCADE
);

-- สร้างตาราง product_images
CREATE TABLE IF NOT EXISTS product_images (
    image_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    product_id UUID NOT NULL,
    image_url VARCHAR(255) NOT NULL,
    alt_text VARCHAR(255),
    is_primary BOOLEAN DEFAULT FALSE,
    sort_order INTEGER DEFAULT 0,
    uploaded_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (product_id) REFERENCES products(product_id) ON DELETE CASCADE
);

-- สร้างตาราง product_options
CREATE TABLE IF NOT EXISTS product_options (
    option_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    product_id UUID NOT NULL,
    name VARCHAR(255) NOT NULL,
    values JSONB NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (product_id) REFERENCES products(product_id) ON DELETE CASCADE
);

-- สร้างตาราง inventory
CREATE TABLE IF NOT EXISTS inventory (
    product_id UUID PRIMARY KEY,
    quantity INTEGER NOT NULL CHECK (quantity >= 0),
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (product_id) REFERENCES products(product_id) ON DELETE CASCADE
);

-- สร้างตาราง shopping_cart
CREATE TABLE IF NOT EXISTS shopping_cart (
    cart_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    status VARCHAR(50) DEFAULT 'active'
);

-- สร้างตาราง shopping_cart_items
CREATE TABLE IF NOT EXISTS shopping_cart_items (
    item_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    cart_id UUID NOT NULL,
    product_id UUID NOT NULL,
    quantity INTEGER NOT NULL CHECK (quantity > 0),
    price NUMERIC(10, 2) NOT NULL CHECK (price >= 0),
    added_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (cart_id) REFERENCES shopping_cart(cart_id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products(product_id) ON DELETE CASCADE
);

-- สร้างตาราง contract
CREATE TABLE IF NOT EXISTS contract (
    contract_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL,
    email VARCHAR(255) NOT NULL,
    phone_number VARCHAR(20),
    message TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- สร้างตาราง users_local
CREATE TABLE IF NOT EXISTS users_local (
    user_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    username VARCHAR(100) UNIQUE NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    provider VARCHAR(50) DEFAULT 'local'
);


-- เพิ่มข้อมูลตาราง categories
INSERT INTO
    categories (name, description,parent_category_id)
VALUES
    ('หนังสือ', 'หนังสือและสื่อสิ่งพิมพ์', NULL),
    ('ดินสอ/ปากกา', 'อุปกรณ์เครื่องเขียน', NULL),
    ('กระดาษ', 'สื่อวาดภาพระบายสี', NULL),
    ('สมุดวาดเขียน', 'จดบันทึกและการวาดเขียน', NULL),
    ('DIY', 'อุปกรณ์งานช่างและการประดิษฐ์', NULL);

-- เพิ่มข้อมูลตาราง sellers -- สร้างเพิ่มได้
INSERT INTO
    sellers (seller_id, name, contact_info)
VALUES
    (
        'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11',
        'บริษัท ไทยเทคโนโลยี จำกัด',
        'contact@thaitechnology.com'
    ),
    (
        'b0eebc99-9c0b-4ef8-bb6d-6bb9bd380a22',
        'สำนักพิมพ์ ก้าวหน้า',
        'info@kaownapublishing.com'
    ),
    (
        'c0eebc99-9c0b-4ef8-bb6d-6bb9bd380a33',
        'แฟชั่นไทยดีไซน์',
        'support@thaifashiondesign.com'
    );


--docker-compose down -v
--docker-compose up --build -d

INSERT INTO products 
(
    name, description, brand, model_number, sku, price, status, seller_id, product_type, category_id
)
VALUES 
(
    'กระดาษA4 Double A', 
    'กระดาษA4 Double A 80g(80แกรม) สี:ขาว   ขนาด:210x297มม.', 
    'กระดาษA4', 
    'A4-001', 
    'SKU-1', 
       125, 
    'active',
    'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 
         'กระดาษ',
          3 
 
);



INSERT INTO products 
(
    name, description, brand, model_number, sku, price, status, seller_id, product_type, category_id
)
VALUES 
(
    'กระดาษลูกฟูก', 
    'กระดาษลูกฟูก 5 ชั้นลอน โดย:Hong Thai Group  ขนาด:40×48นิ้ว', 
    'Hong Thai Group', 
    'sa-001', 
    'SKU-2', 
    58, 
    'active', 
    'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 
    'กระดาษ', 
     3
);



INSERT INTO products 
(
    name, description, brand, model_number, sku, price, status, seller_id, product_type, category_id
)
VALUES 
(
    'กระดาษร้อยปอนด์', 
    'กระดาษร้อยปอนด์ 200g(200แกรม) สี:ขาว ขนาด:29x42ซม.', 
    'Hong Thai Group', 
    'fgdh-001', 
    'SKU-3', 
    12, 
    'active', 
    'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 
    'กระดาษ', 
    3
);


INSERT INTO products 
(
    name, description, brand, model_number, sku, price, status, seller_id, product_type, category_id
)
VALUES 
(
    'กระดาษฟอกขาว', 
    'กระดาษฟอกขาว Woodfree Inkjet paper roll กระดาษม้วนสี:ขาว', 
    'SDD', 
    'ghf-001', 
    'SKU-4', 
    239, 
    'active', 
    'b0eebc99-9c0b-4ef8-bb6d-6bb9bd380a22', 
    'กระดาษ', 
    3
);


INSERT INTO products 
(
    name, description, brand, model_number, sku, price, status, seller_id, product_type, category_id
)
VALUES 
(
    'กระดาษPVC', 
    'กระดาษPVC โดย DTAWAN  135g (135แกรม) ขาวขุ่น ขนาด 12x18นิ้ว', 
    'DTAWAN', 
    'FDS-23', 
    'SKU-5', 
    1000, 
    'active', 
    'c0eebc99-9c0b-4ef8-bb6d-6bb9bd380a33', 
    'กระดาษ', 
    3
);





INSERT INTO products 
(
    name, description, brand, model_number, sku, price, status, seller_id, product_type, category_id
)
VALUES 
(
    'ปากกาลูกลื่น', 
    'ปากกาลูกลื่น Quantum รุ่น Moji 0.29มม. [กล่อง 12 ด้าม]', 
    'PENTEL', 
    'FDS-56', 
    'SKU-6', 
    168, 
    'active', 
    'b0eebc99-9c0b-4ef8-bb6d-6bb9bd380a22', 
    'ดินสอ/ปากกา', 
    2
);


INSERT INTO products 
(
    name, description, brand, model_number, sku, price, status, seller_id, product_type, category_id
)
VALUES 
(
    'ปากกาเจล', 
    'ปากกาเจล Pentel Energel X รุ่น BLN105 BL107 และ ไส้ปากกา 0.5 0.7 MM', 
    'PENTEL', 
    'DFS236', 
    'SKU-7', 
    30, 
    'active', 
    'b0eebc99-9c0b-4ef8-bb6d-6bb9bd380a22', 
    'ดินสอ/ปากกา', 
    2
);


INSERT INTO products 
(
    name, description, brand, model_number, sku, price, status, seller_id, product_type, category_id
)
VALUES 
(
    'Master Art', 
    'Master Art ดินสอ ดินสอดำ HB อาเซียนซีรี่ย์ จำนวน 12 แท่ง  ราคา 36', 
    'Master Art', 
    'DFS236', 
    'SKU-8', 
    36, 
    'active', 
    'c0eebc99-9c0b-4ef8-bb6d-6bb9bd380a33', 
    'ดินสอ/ปากกา', 
    2
);



INSERT INTO products 
(
    name, description, brand, model_number, sku, price, status, seller_id, product_type, category_id
)
VALUES 
(
    'ดินสอกด Quantum', 
    'Metal Mechanism กลไกโลหะคุณภาพสูง สามารถจับไส้ได้มั่นคงเพร้อมทั้งควบคุมการไฟลของไส้ได้ยอดเยี่ยม', 
    'Quantum', 
    'DFS236', 
    'SKU-9', 
    20, 
    'active', 
    'c0eebc99-9c0b-4ef8-bb6d-6bb9bd380a33', 
    'ดินสอ/ปากกา', 
    2
);


INSERT INTO products 
(
    name, description, brand, model_number, sku, price, status, seller_id, product_type, category_id
)
VALUES 
(
    'ดินสอกด Pentel', 
    'Metal Mechanism กลไกโลหะคุณภาพสูง สามารถจับไส้ได้มั่นคงเพร้อมทั้งควบคุมการไฟลของไส้ได้ยอดเยี่ยม', 
    'Pentel', 
    'FVCS', 
    'SKU-10', 
    25, 
    'active', 
    'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 
    'ดินสอ/ปากกา', 
    2
);


INSERT INTO products 
(
    name, description, brand, model_number, sku, price, status, seller_id, product_type, category_id
)
VALUES 
(
    'Renaissance(300)', 
    'Renaissance (เรนาซองซ์) สมุดวาดเขียน 300 แกรม แบบหยาบ ผลิตจากเยื่อกระดาษชั้นดีแบบเยื่อใยขาว ทำให้การกระจายตัวของสีเป็นวงกว้าง เนื้อกระดาษซับน้ำได้ดี ไม่ทำให้สีดูจืด ดูดซับสีและเก็บเนื้อสีได้ดี มีลาย Texture บนผิว ปราศจากกรด (Acid Free) ทำให้กระดาษเก็บไว้ได้นาน ไม่ซีดเหลือง', 
    'Renaissance', 
    'FSC-12', 
    'SKU-15', 
    153, 
    'active', 
    'b0eebc99-9c0b-4ef8-bb6d-6bb9bd380a22', 
    'สมุดวาดเขียน', 
    4
);


INSERT INTO products 
(
    name, description, brand, model_number, sku, price, status, seller_id, product_type, category_id
)
VALUES 
(
    'Renaissance (เรนาซองซ์)', 
    'Renaissance (เรนาซองซ์) สมุดวาดเขียน 200 แกรม แบบหยาบ ผลิตจากเยื่อกระดาษชั้นดี', 
    'Renaissance', 
    'JHD', 
    'SKU-11', 
    153, 
    'active', 
    'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 
    'สมุดวาดเขียน', 
    4
);


INSERT INTO products 
(
    name, description, brand, model_number, sku, price, status, seller_id, product_type, category_id
)
VALUES 
(
    'Canson สมุดวาดเขียนสีน้ำ', 
    'Canson สมุดวาดเขียนสีน้ำ A5 300 แกรม ชนิดเรียบเข้าเล่มแบบสันกาว เปิดกางออกสะดวก ฉีกใช้งานง่าย เนื้อในกระดาษ 100 ปอนด์ ผิวเรียบ ไม่เหลือง แม้เก็บนาน เหมาะสำหรับงานออกแบบ', 
    'Canson', 
    'DSC', 
    'SKU-12', 
    106, 
    'active', 
    'c0eebc99-9c0b-4ef8-bb6d-6bb9bd380a33', 
    'สมุดวาดเขียน', 
    4
);


INSERT INTO products 
(
    name, description, brand, model_number, sku, price, status, seller_id, product_type, category_id
)
VALUES 
(
    'Master Art (มาสเตอร์อาร์ต) สมุดวาดเขียน', 
    'Master Art (มาสเตอร์อาร์ต) สมุดวาดเขียน สมุดวาดภาพS-104 ร้อยลวดเข้าเล่มแบบสันห่วง เหมาะสำหรับงานเขียนด้วยปากกาเขียนแบบ ,หัวสักหลาด ,หัวเข็ม และ ดินสอวาดเขียน', 
    'Master Art', 
    'FHD', 
    'SKU-13', 
     91, 
    'active', 
    'b0eebc99-9c0b-4ef8-bb6d-6bb9bd380a22', 
    'สมุดวาดเขียน', 
    4
);



INSERT INTO products 
(
    name, description, brand, model_number, sku, price, status, seller_id, product_type, category_id
)
VALUES 
(
    'จิตวิทยาสายดาร์ก', 
    'เทคนิคทางจิตวิทยาที่ช่วยให้คุณใช้คำพูดควบคุมจิตใจคนทำให้พวกเขาคล้อยตามและทำอย่างที่คุณต้องการโดยไม่รู้ตัว', 
    'WeLearn', 
    '34EW', 
    'SKU-16', 
    250, 
    'active', 
    'b0eebc99-9c0b-4ef8-bb6d-6bb9bd380a22', 
    'หนังสือ', 
    1
);


INSERT INTO products 
(
    name, description, brand, model_number, sku, price, status, seller_id, product_type, category_id
)
VALUES 
(
    'ร้านขายยาปริศนารับแก้ปัญหาหัวใจ', 
    'ในเขตที่กำลังจะพัฒนาใหม่เขตหนึ่งของกรุงโซลซึ่งเต็มไปด้วยบ้านเก่าที่พังทลาย มีร้านขายยาน่าสงสัยร้านหนึ่งซึ่งแตกต่างจากร้านอื่น ๆ เพราะมีทั้งป้ายชื่อว่า “ร้านยารัก” มีทั้งเสียงดนตรี กลิ่นหอมของชาสมุนไพร มีเก้าอี้นวมที่ให้คำปรึกษา', 
    'Piccolo', 
    'JG562', 
    'SKU-17', 
    385, 
    'active', 
    'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 
    'หนังสือ', 
    1
);


INSERT INTO products 
(
    name, description, brand, model_number, sku, price, status, seller_id, product_type, category_id
)
VALUES 
(
    'เถียงอย่างไรให้ชนะแมว (How to Argue with a Cat)', 
    'ศิลปะการพูดที่จะทำให้คุณ ครองใจคนทั้งโลกได้', 
    'Heinrichs', 
    'DKF54', 
    'SKU-18', 
    175, 
    'active', 
    'c0eebc99-9c0b-4ef8-bb6d-6bb9bd380a33', 
    'หนังสือ', 
    1
);


INSERT INTO products 
(
    name, description, brand, model_number, sku, price, status, seller_id, product_type, category_id
)
VALUES 
(
    'Practical DEVOPS and Cloud Engineering', 
    'หนังสือเล่มนี้เป็นคู่มือพัฒนาทักษะทางด้าน Practical DEVOPS and Cloud Engineering ภาคปฏิบัติ มีจุดประสงค์เพื่อพัฒนาทักษะ 2 ด้าน คือ Reskill และ Upskill เหมาะสำหรับนักศึกษา โปรแกรมเมอร์ นักพัฒนาซอฟต์แวร์และผู้ที่อยากข้ามสายงานมาทำงานทางด้านนี้', 
    'Infopress', 
    'GHDE23', 
    'SKU-19', 
    545, 
    'active', 
    'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 
    'หนังสือ', 
    1
);


INSERT INTO products 
(
    name, description, brand, model_number, sku, price, status, seller_id, product_type, category_id
)
VALUES 
(
    'TBX บุกทะลวงข้อสอบ TOEIC LISTENING READING 1000 ข้อ', 
    'ฝึกให้ชัวร์ด้วยข้อสอบ TOEIC เสมือนจริงเหมาะทั้งเรียนรู้ด้วยตัวเอง หรือติวเตอร์ใช้ประกอบการสอนข้อสอบ TOEIC รูปแบบใหม่ เสมือนจริง 5 ชุด 1000 ข้อ จุใจ รวมข้อสอบทั้งพาร์ Listening + Readingเหมาะสำหรับผู้สอบที่มุ่งเป๊าคะแนน 500+ถึง 800+', 
    'Think Beyond', 
    'HGD12', 
    'SKU-20', 
    750, 
    'active', 
    'b0eebc99-9c0b-4ef8-bb6d-6bb9bd380a22', 
    'หนังสือ', 
    1
);



INSERT INTO products 
(
    name, description, brand, model_number, sku, price, status, seller_id, product_type, category_id
)
VALUES 
(
    'MONT MARTE เฟรมผ้าใบ', 
    'MONT MARTE เฟรมผ้าใบ รุ่น Double Thick (D.T.) สีขาว ขนาด 35.6x71.1 ซม.', 
    'MONT MARTE', 
    'HDYR2', 
    'SKU-21', 
    735, 
    'active', 
    'c0eebc99-9c0b-4ef8-bb6d-6bb9bd380a33', 
    'DIY', 
    5
);


INSERT INTO products 
(
    name, description, brand, model_number, sku, price, status, seller_id, product_type, category_id
)
VALUES 
(
    'เฟรมผ้าใบระบายสีวงกลม 30 ซม. Montmarte คละลาย', 
    'เฟรมผ้าใบระบายสีวงกลม 30 ซม. Montmarte คละลาย ผ้าใบหนา 380 แกรม ช่วยป้องกันไม่ให้สีซึมลงบนพื้นผิว', 
    'MONT MARTE', 
    'HDYR2', 
    'SKU-22', 
    415, 
    'active', 
    'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 
    'DIY', 
    5
);


INSERT INTO products 
(
    name, description, brand, model_number, sku, price, status, seller_id, product_type, category_id
)
VALUES 
(
    'เชือกขาว', 
    'เชือกขาว 30 เส้น 13 หลา บอสตัน เชือกยาว 30 หลา', 
    'CGS', 
    'HYGD78', 
    'SKU-23', 
    25, 
    'active', 
    'b0eebc99-9c0b-4ef8-bb6d-6bb9bd380a22', 
    'DIY', 
    5

);


INSERT INTO products 
(
    name, description, brand, model_number, sku, price, status, seller_id, product_type, category_id
)
VALUES 
(
    'ใบมีดไทเทเนียม', 
    'ใบมีดไทเทเนียม 3M Size L Scotch XP002015301 คลือบไทเทเนียมเพิ่มความแข็งแรงและคงทน จึงคมยาวนานกว่าถึง 2 เท่า', 
    'Scotch', 
    'XP002015301', 
    'SKU-24', 
    150, 
    'active', 
    'c0eebc99-9c0b-4ef8-bb6d-6bb9bd380a33', 
    'DIY', 
    5
);


INSERT INTO products 
(
    name, description, brand, model_number, sku, price, status, seller_id, product_type, category_id
)
VALUES 
(
    'เครื่องตัดอเนกประสงค์', 
    'เครื่องตัดอเนกประสงค์ BROTHER รุ่น SDX1200 สามารถตัดวัสดุความหนาได้ถึง 3 มิลลิเมตร', 
    'BROTHER', 
    'SDX1200', 
    'SKU-25', 
    22900, 
    'active', 
    'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 
    'DIY', 
    5
);


------------------------------------------------------------------------------

INSERT INTO inventory 
(
    product_id, quantity
)
VALUES 
(
    (SELECT product_id FROM products WHERE sku = 'SKU-1'), 100
);

INSERT INTO inventory 
(
    product_id, quantity
)
VALUES 
(
    (SELECT product_id FROM products WHERE sku = 'SKU-2'), 200
);

INSERT INTO inventory 
(
    product_id, quantity
)
VALUES 
(
    (SELECT product_id FROM products WHERE sku = 'SKU-3'), 150
);

INSERT INTO inventory 
(
    product_id, quantity
)
VALUES 
(
    (SELECT product_id FROM products WHERE sku = 'SKU-4'), 50
);

INSERT INTO inventory 
(
    product_id, quantity
)
VALUES 
(
    (SELECT product_id FROM products WHERE sku = 'SKU-5'), 300
);

INSERT INTO inventory 
(
    product_id, quantity
)
VALUES 
(
    (SELECT product_id FROM products WHERE sku = 'SKU-6'), 150
);

INSERT INTO inventory 
(
    product_id, quantity
)
VALUES 
(
    (SELECT product_id FROM products WHERE sku = 'SKU-7'), 100
);

INSERT INTO inventory 
(
    product_id, quantity
)
VALUES 
(
    (SELECT product_id FROM products WHERE sku = 'SKU-8'), 120
);

INSERT INTO inventory 
(
    product_id, quantity
)
VALUES 
(
    (SELECT product_id FROM products WHERE sku = 'SKU-9'), 200
);

INSERT INTO inventory 
(
    product_id, quantity
)
VALUES 
(
    (SELECT product_id FROM products WHERE sku = 'SKU-10'), 100
);

INSERT INTO inventory 
(
    product_id, quantity
)
VALUES 
(
    (SELECT product_id FROM products WHERE sku = 'SKU-11'), 80
);

INSERT INTO inventory 
(
    product_id, quantity
)
VALUES 
(
    (SELECT product_id FROM products WHERE sku = 'SKU-12'), 60
);

INSERT INTO inventory 
(
    product_id, quantity
)
VALUES 
(
    (SELECT product_id FROM products WHERE sku = 'SKU-13'), 70
);

INSERT INTO inventory 
(
    product_id, quantity
)
VALUES 
(
    (SELECT product_id FROM products WHERE sku = 'SKU-15'), 100
);

INSERT INTO inventory 
(
    product_id, quantity
)
VALUES 
(
    (SELECT product_id FROM products WHERE sku = 'SKU-16'), 200
);

INSERT INTO inventory 
(
    product_id, quantity
)
VALUES 
(
    (SELECT product_id FROM products WHERE sku = 'SKU-17'), 80
);

INSERT INTO inventory 
(
    product_id, quantity
)
VALUES 
(
    (SELECT product_id FROM products WHERE sku = 'SKU-18'), 90
);

INSERT INTO inventory 
(
    product_id, quantity
)
VALUES 
(
    (SELECT product_id FROM products WHERE sku = 'SKU-19'), 60
);

INSERT INTO inventory 
(
    product_id, quantity
)
VALUES 
(
    (SELECT product_id FROM products WHERE sku = 'SKU-20'), 150
);

INSERT INTO inventory 
(
    product_id, quantity
)
VALUES 
(
    (SELECT product_id FROM products WHERE sku = 'SKU-21'), 200
);

INSERT INTO inventory 
(
    product_id, quantity
)
VALUES 
(
    (SELECT product_id FROM products WHERE sku = 'SKU-22'), 180
);

INSERT INTO inventory 
(
    product_id, quantity
)
VALUES 
(
    (SELECT product_id FROM products WHERE sku = 'SKU-23'), 220
);

INSERT INTO inventory 
(
    product_id, quantity
)
VALUES 
(
    (SELECT product_id FROM products WHERE sku = 'SKU-24'), 90
);

INSERT INTO inventory 
(
    product_id, quantity
)
VALUES 
(
    (SELECT product_id FROM products WHERE sku = 'SKU-25'), 160
);

INSERT INTO product_images (product_id, image_url, is_primary, sort_order) 
VALUES 
    ((SELECT product_id FROM products WHERE sku = 'SKU-1'), 'https://delivery.doubleapaper.com/wp-content/uploads/2018/04/208140800312652-2.png', TRUE, 1),
    ((SELECT product_id FROM products WHERE sku = 'SKU-2'), 'https://hongthaipackaging.com/wp-content/uploads/2019/05/%E0%B9%81%E0%B8%9C%E0%B8%99%E0%B8%81%E0%B8%A3%E0%B8%B0%E0%B8%94%E0%B8%B2%E0%B8%A9%E0%B8%A5%E0%B8%B9%E0%B8%81%E0%B8%9F%E0%B8%B9%E0%B8%81-%E0%B8%A5%E0%B8%AD%E0%B8%99-E.jpg', TRUE, 2),
    ((SELECT product_id FROM products WHERE sku = 'SKU-3'), 'https://www.dohome.co.th/media/catalog/product/cache/e446f15aaa8dc66b80b7a0df334f7c5a/1/0/10357823_rem_1200_2.jpg', TRUE, 3),
    ((SELECT product_id FROM products WHERE sku = 'SKU-4'), 'https://down-th.img.susercontent.com/file/9a499681851180f3c59412d9316b78c3', TRUE, 4),
    ((SELECT product_id FROM products WHERE sku = 'SKU-5'), 'https://dm.lnwfile.com/_webp_resize_images/300/300/r2/69/2s.webp', TRUE, 5),
    ((SELECT product_id FROM products WHERE sku = 'SKU-6'), 'https://down-th.img.susercontent.com/file/58a33ed89ebf18126432d5c0308ea71d', TRUE, 6),
    ((SELECT product_id FROM products WHERE sku = 'SKU-7'), 'https://down-th.img.susercontent.com/file/9dccbe995d4a177e61bad095b061ac2b', TRUE, 7),
    ((SELECT product_id FROM products WHERE sku = 'SKU-8'), 'https://www.siriwongpanid.com/wp-content/uploads/2021/07/1011012.jpg', TRUE, 8),
    ((SELECT product_id FROM products WHERE sku = 'SKU-9'), 'https://officebulkydhas.vtexassets.com/arquivos/ids/165866/253950_quantum_mechanical_pencil_0.5_mm_atom_qm239_0.5_assorted_colors_02.jpg?v=638155294434030000', TRUE, 9),
    ((SELECT product_id FROM products WHERE sku = 'SKU-10'), 'https://inwfile.com/s-cp/9qx3kh.jpg', TRUE, 10),
    ((SELECT product_id FROM products WHERE sku = 'SKU-11'), 'https://down-th.img.susercontent.com/file/da08e993c72906ed2d62618aa510d1bd', TRUE, 11),
    ((SELECT product_id FROM products WHERE sku = 'SKU-12'), 'https://image.makewebeasy.net/makeweb/0/lZqXXO3KV/DefaultData/f0966f01bfda8ceede36ec29ae9c2263.jpg?v=202405291424', TRUE, 12),
    ((SELECT product_id FROM products WHERE sku = 'SKU-13'), 'https://down-th.img.susercontent.com/file/5368c068415968cfa92a8d1f7217a72f', TRUE, 13),
    ((SELECT product_id FROM products WHERE sku = 'SKU-15'), 'https://down-th.img.susercontent.com/file/24063e37d4138c65836815c2cd39340f', TRUE, 14),
    ((SELECT product_id FROM products WHERE sku = 'SKU-16'), 'https://api.chulabook.com/images/pid-176733.jpg', TRUE, 15),
    ((SELECT product_id FROM products WHERE sku = 'SKU-17'), 'https://down-th.img.susercontent.com/file/th-11134207-7rasc-m0sw7t114gcac8', TRUE, 16),
    ((SELECT product_id FROM products WHERE sku = 'SKU-18'), 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRFaTuOMVMs5NuNecsk-drIjpptmkW7N6KWWw&s', TRUE, 17),
    ((SELECT product_id FROM products WHERE sku = 'SKU-19'), 'https://img.lazcdn.com/g/p/c76486e7f7fbd4d93b2b54f867eb1201.jpg_720x720q80.jpg', TRUE, 18),
    ((SELECT product_id FROM products WHERE sku = 'SKU-20'), 'https://api.chulabook.com/images/pid-169070.jpg', TRUE, 19),
    ((SELECT product_id FROM products WHERE sku = 'SKU-21'), 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSO5w7izP9dQ0RvGB8eCoVobvy4uBoXtGGEhg&s', TRUE, 20),
    ((SELECT product_id FROM products WHERE sku = 'SKU-22'), 'https://pim-cdn0.ofm.co.th/products/large/K091839.jpg', TRUE, 21),
    ((SELECT product_id FROM products WHERE sku = 'SKU-23'), 'https://down-th.img.susercontent.com/file/ea086020b275c2b26e224a38ecd3b48c', TRUE, 22),
    ((SELECT product_id FROM products WHERE sku = 'SKU-24'), 'https://down-th.img.susercontent.com/file/83ce35d311194946ff327dba97591bfb', TRUE, 23),
    ((SELECT product_id FROM products WHERE sku = 'SKU-25'), 'https://down-th.img.susercontent.com/file/th-11134207-7r98p-ls7z8btfd2gt89', TRUE, 24);


INSERT INTO users_local (email, password, username) 
VALUES ('katchapong_j@silpakorn.edu', 'password', 'test_jakkaphat');

INSERT INTO shopping_cart (user_id)
VALUES ((SELECT user_id FROM users_local WHERE email = 'katchapong_j@silpakorn.edu'));