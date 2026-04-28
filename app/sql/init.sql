DROP TABLE IF EXISTS tasks;

CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    completed BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO tasks (title, description, completed) VALUES
    ('Купить продукты', 'Молоко, хлеб, яйца, масло', false),
    ('Сделать уборку', 'Помыть полы, вытереть пыль, пропылесосить', false),
    ('Позвонить маме', 'Спросить про здоровье и планы на выходные', false),
    ('Закончить отчет', 'Подготовить квартальный отчет для начальника', true),
    ('Сходить в спортзал', 'Пробежка 5 км и тренировка спины', false),
    ('Полить цветы', 'Все комнатные растения нуждаются в поливе', true),
    ('Заплатить за интернет', 'Оплатить до 10 числа', false),
    ('Почитать книгу', 'Прочитать 50 страниц новой книги', false),
    ('Записаться к врачу', 'Профилактический осмотр', true),
    ('Помыть машину', 'Заехать на мойку после работы', false);