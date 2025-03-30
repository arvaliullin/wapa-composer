
CREATE TABLE experiment_design (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    config JSONB NOT NULL,
    js_path TEXT,
    wasm_path TEXT NOT NULL
);

COMMENT ON TABLE experiment_design IS 'Таблица для хранения информации о конфигурации экспериментов';
COMMENT ON COLUMN experiment_design.id IS 'Уникальный идентификатор эксперимента';
COMMENT ON COLUMN experiment_design.name IS 'Название эксперимента';
COMMENT ON COLUMN experiment_design.config IS 'Конфигурация эксперимента';
COMMENT ON COLUMN experiment_design.js_path IS 'Путь к JavaScript файлу';
COMMENT ON COLUMN experiment_design.wasm_path IS 'Путь к WebAssembly файлу';


CREATE TABLE experiment_execution (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    experiment_id UUID REFERENCES experiment_design(id) ON DELETE CASCADE,
    execution_time BIGINT NOT NULL,
    repetitions INTEGER NOT NULL CHECK (repetitions > 0)
);


COMMENT ON TABLE experiment_execution IS 'Таблица для хранения информации о выполнении экспериментов';
COMMENT ON COLUMN experiment_execution.id IS 'Уникальный идентификатор выполнения';
COMMENT ON COLUMN experiment_execution.experiment_id IS 'Конфигурация эксперимента из таблицы';
COMMENT ON COLUMN experiment_execution.execution_time IS 'Время выполнения кода';
COMMENT ON COLUMN experiment_execution.repetitions IS 'Количество повторений выполнения';
