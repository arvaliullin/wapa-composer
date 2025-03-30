CREATE SCHEMA IF NOT EXISTS composer;

CREATE EXTENSION IF NOT EXISTS pgcrypto SCHEMA composer;

CREATE TABLE composer.design (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    description JSONB NOT NULL,
    js_filename TEXT,
    wasm_filename TEXT NOT NULL
);

COMMENT ON TABLE composer.design IS 'Таблица для хранения плана эксперимента';
COMMENT ON COLUMN composer.design.id IS 'Уникальный идентификатор эксперимента';
COMMENT ON COLUMN composer.design.name IS 'Название эксперимента';
COMMENT ON COLUMN composer.design.description IS 'Описание эксперимента';
COMMENT ON COLUMN composer.design.js_filename IS 'JavaScript файл';
COMMENT ON COLUMN composer.design.wasm_filename IS 'WebAssembly файл';


CREATE TABLE composer.experiment (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    design_id UUID REFERENCES composer.design(id) ON DELETE CASCADE,
    execution_time BIGINT NOT NULL,
    repetitions INTEGER NOT NULL CHECK (repetitions > 0)
);

COMMENT ON TABLE composer.experiment IS 'Таблица для хранения информации о выполнении экспериментов';
COMMENT ON COLUMN composer.experiment.id IS 'Уникальный идентификатор эксперимента';
COMMENT ON COLUMN composer.experiment.design_id IS 'План эксперимента';
COMMENT ON COLUMN composer.experiment.execution_time IS 'Время выполнения кода';
COMMENT ON COLUMN composer.experiment.repetitions IS 'Количество повторений выполнения';
