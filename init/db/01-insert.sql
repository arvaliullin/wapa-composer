INSERT INTO composer.design (name, description, js_filename, wasm_filename)
VALUES 
    ('Experiment 1', '{
    "name": "x2Integrate",
    "lang": "go",
    "functions": [
        {
            "function": "x2Integrate",
            "args": [
                0.0,
                100.0,
                10000
            ]
        },
        {
            "function": "x2Integrate",
            "args": [
                0.0,
                50.0,
                5000
            ]
        }
    ]
}', 'experiment1.js', 'experiment1.wasm'),
    ('Experiment 2', '{"param": "value2"}', 'experiment2.js', 'experiment2.wasm');

INSERT INTO composer.experiment (design_id, execution_time, repetitions)
VALUES 
    ((SELECT id FROM composer.design WHERE name = 'Experiment 1'), 123456, 10),
    ((SELECT id FROM composer.design WHERE name = 'Experiment 2'), 789012, 5);
