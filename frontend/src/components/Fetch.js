import React, { Component } from 'react';

class JsonFetcher extends Component {
    constructor() {
        super();
        this.state = {
            data: null,
            error: null,
        };
    }

    componentDidMount() {
        // Здесь вы можете выполнить ваш POST-запрос или другой тип запроса
        fetch("http://localhost:8080/api/branches", {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json', // Устанавливаем заголовок для JSON-данных
            },
            body: JSON.stringify({
                "start": {
                    "latitude": 37.7749,
                    "longitude": -122.4194
                },
                "end": {
                    "latitude": 37.3352,
                    "longitude": -121.8811
                },
                "transportType": "car"
            }), // Отправляем JSON-данные
        })
            .then(response => {
                if (!response.ok) {
                    throw new Error('Ошибка сети'); // Обрабатываем ошибки сети
                }
                return response.json(); // Разбираем JSON-ответ
            })
            .then(data => {
                console.log(data); // Выводим данные в консоль
            })
            .catch(error => {
                console.error('Произошла ошибка:', error); // Обрабатываем ошибки
            });
    }

    render() {
        const { data, error } = this.state;

        if (error) {
            return <div>Произошла ошибка: {error.message}</div>;
        }

        if (data) {
            return (
                <div>
                    <h2>JSON-данные:</h2>
                    <pre>{JSON.stringify(data, null, 2)}</pre>
                </div>
            );
        }

        return <div>Загрузка...</div>;
    }
}

export default JsonFetcher;
