import {YMaps, Map, Placemark, Polyline} from '@pbe/react-yandex-maps';
import "../../App.css";
const Direction = (props) => (
    <>

        <Polyline
            geometry={props.coord}
            options={{
                balloonCloseButton: false,
                strokeColor: "#000",
                strokeWidth: 4,
                strokeOpacity: 0.5,
            }}
        />

    </>
);

export default Direction;
