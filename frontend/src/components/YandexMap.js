import {YMaps, Map, Placemark, GeolocationControl, Button, Polyline, ZoomControl} from '@pbe/react-yandex-maps';
import styles from "./YandexMap.css"
import Atms from "./Atms";

const YandexMap = (props) => (
    <>

    <YMaps>
        <div>
            <Map width={window.innerWidth} height={window.innerHeight-180}//это карта
                 defaultState={{ center: [ 55.75, 37.57 ], zoom: 9 }}>
                <Atms data = {props.data} filters/>
                <ZoomControl options={{ float: "left" }} />
                <Placemark geometry={props.me}
                           options={{
                               preset: "islands#RedCircleDotIcon",
                               groupByCoordinates: false,

                           }}/>
                {/*<Polyline //эта штука отрисовывает линию*/}
                {/*    geometry={props.coordinates}*/}
                {/*    options={{*/}
                {/*        balloonCloseButton: false,*/}
                {/*        strokeColor: "#000",*/}
                {/*        strokeWidth: 4,*/}
                {/*        strokeOpacity: 0.5,*/}
                {/*    }}*/}
                {/*/>*/}
            </Map>
        </div>
    </YMaps>
    </>
);

export default YandexMap;
