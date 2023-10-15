import * as React from 'react';
import ReactDOM from 'react-dom';
import App from "./App"
import {
    AdaptivityProvider,
    ConfigProvider,

} from '@vkontakte/vkui';
import '@vkontakte/vkui/dist/vkui.css';



ReactDOM.render(
    <ConfigProvider>
        <AdaptivityProvider>
            <App />
        </AdaptivityProvider>
    </ConfigProvider>,
    document.getElementById('root'),
);
