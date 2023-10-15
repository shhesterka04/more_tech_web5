import {YMaps, Map, Placemark, Clusterer} from '@pbe/react-yandex-maps';
import "../App.css";
const Atms = (props) => (
                <div>
                    <Clusterer
                        options={{
                            preset: "islands#blueClusterIcons",
                            groupByCoordinates: false,
                        }}
                    >
                        {props.data.map((item) => (
                            <Placemark geometry={[item.latitude, item.longitude]}
                                       options={{
                                           preset: "islands#darkBlueCircleDotIcon",
                                           groupByCoordinates: false,
                                           click: console.log(item.address)
                                       }}

                                        />
                        ))}
                    </Clusterer>

                    {/*<Placemark geometry={[55.684758, 37.738521]}*/}
                    {/*    options={{*/}
                    {/*        iconColor:"black",*/}
                    {/*        presetStorage:"islands#dotIcon"*/}
                    {/*    }}*/}
                    {/*/>*/}
                </div>
);

export default Atms;
