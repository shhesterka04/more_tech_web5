import * as React from 'react';
import ReactDOM from 'react-dom';
import axios from 'axios';
import styles from "./App.css"
import ATMS from './json/atms.json';
import CORD from './json/direction.json';
import { Icon24Filter } from '@vkontakte/icons';
import {
  AdaptivityProvider,
  ConfigProvider,
  AppRoot,
  SplitLayout,
  SplitCol,
  View,
  Panel,
  PanelHeader,
  Header,
  Group,
  SimpleCell,
  ModalRoot,
  ModalPage,
  PanelHeaderClose,
  ModalPageHeader,
  PanelHeaderSubmit,
  CellButton,
  FormItem,
  Button,
  Radio,
  SelectMimicry,
  Checkbox,
  SegmentedControl,
  Separator,
  Div,
  Input, usePlatform, useAdaptivityConditionalRender, useAdaptivityWithJSMediaQueries, useAdaptivity
} from '@vkontakte/vkui';
import '@vkontakte/vkui/dist/vkui.css';
import {useEffect, useState} from "react";
import YandexMap from "./components/YandexMap";
import JsonFetcher from "./components/Fetch";

function myGeolocation() {
  if (navigator.geolocation) {
    navigator.geolocation.getCurrentPosition(
        function (position) {
          // Обработка успешного получения местоположения
          const latitude = position.coords.latitude;
          const longitude = position.coords.longitude;
          console.log(`Широта: ${latitude}, Долгота: ${longitude}`);
          return [latitude, longitude];
        },
        function (error) {
          // Обработка ошибки
          console.error(`Ошибка при получении местоположения: ${error.message}`);
        }
    );
  } else {
    console.error('Геолокация не поддерживается вашим браузером.');
  }
}

const App = () => {


  const { sizeX, sizeY } = useAdaptivity();
  const { isDesktop } = useAdaptivityWithJSMediaQueries();
  const [activeModal, setActiveModal] = useState(null);
  const [modalHistory, setModalHistory] = useState([]);
  const [reportType, changeReportType] = React.useState('idea');
  const [selectedType, changeSelectedType] = React.useState();
  const [data, setData] = useState([]);
  const [coord, setCoord] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const platform = usePlatform();


  const [isOffice, setIsOffice] = useState(true);
  const [nfc, setNfc] = useState(true);
  const [qr, setQr] = useState(true);
  const [blind, setBlind] = useState(true);
  const [wheelchair, setWheelchair] = useState(true);
  const [face, setFace] = useState(true);
  const [allday, setAllday] = useState(true);

  const handleIsOfficeChange = () => {
    setIsOffice(!isOffice);
  };

  const handleNfcChange = () => {
    setNfc(!nfc);
  };

  const handleQrChange = () => {
    setQr(!qr);
  };

  const handleBlindChange = () => {
    setBlind(!blind);
  };

  const handleWheelchairChange = () => {
    setWheelchair(!wheelchair);
  };

  const handleFaceChange = () => {
    setFace(!face);
  };

  const handleAlldayChange = () => {
    setAllday(!allday);
  };

  async function fetchData(url) {
    try {
      const response = await axios.get(url);
      if (response.status === 200) {
        return response.data; // Возвращаем JSON-данные
      } else {
        throw new Error(`Ошибка: ${response.status}`);
      }
    } catch (error) {
      console.error('Произошла ошибка:', error);
      throw error; // Можно обработать ошибку в вызывающем коде
    }
  }

  const changeActiveModal = (activeModal) => {
    activeModal = activeModal || null;
    let localModalHistory = modalHistory ? [...modalHistory] : [];

    if (activeModal === null) {
      localModalHistory = [];
    } else if (modalHistory.indexOf(activeModal) !== -1) {
      localModalHistory = localModalHistory.splice(0, localModalHistory.indexOf(activeModal) + 1);
    } else {
      localModalHistory.push(activeModal);
    }

    setActiveModal(activeModal);
    setModalHistory(localModalHistory);
  };

  const modalBack = () => {
    changeActiveModal(modalHistory[modalHistory.length - 2]);
  }
  let MyC = null;

    useEffect(() => {
      MyC = myGeolocation();

      // Здесь вы можете указать URL вашего бэкенда
      const isOffice = true;
      const nfc = true;
      const qr = true;
      const blind = true;
      const wheelchair = true;
      const face = true;
      const allday = true;

      // Получение списка всех отделений по фильтрам:
      //const FilterApiUrl = "/api/branches?isOffice="+"isOffice"+"&nfc=+"nfc}&qr=${qr}&blind=${blind}&wheelchair=${wheelchair}&face=${face}&allday=${allday}";

      // Получение информации о конкретном отделении, включая прогнозируемую нагрузку:
      const IdApiUrl = `http://localhost:80/api/branches/\${branchId}`;

      // Получение информации о оптимальном отделении, включая прогнозируемую нагрузку:
      const SuperApiUrl = `http://localhost:80/api/branches/recommended?location=latitude,longitude&preferredTime=value`;

      fetchData(FilterApiUrl)
          .then((data) => {
            console.log('Полученные данные:', data);
            // Здесь вы можете обработать полученные данные
          })
          .catch((error) => {
            console.error('Произошла ошибка:', error);
            // Здесь вы можете обработать ошибку
          });
    }, []);
  const modal = (
      <ModalRoot activeModal={activeModal}>
        {selectedType ? <ModalPage
            id="filters"
            onClose={modalBack}
            header={
              <ModalPageHeader


              >
                Фильтры отделений
              </ModalPageHeader>
            }
        >
          <Group>



            <FormItem >
              <Radio  name="sex" value={0} defaultChecked>
                Физические лица
              </Radio>
              <Radio name="sex" value={1}>
                Юридические лица
              </Radio>
            </FormItem>
            <Separator />

            <FormItem>
              <Checkbox checked={nfc} onChange={handleNfcChange}>
                NFC для банковских карт
              </Checkbox>
              <Checkbox checked={qr} onChange={handleQrChange}>
                Снятие наличных с помощью QR кода
              </Checkbox>
              <Checkbox checked={isOffice} onChange={handleIsOfficeChange}>
                Зона премиального обслуживания
              </Checkbox>
            </FormItem>
            <Separator />
            <FormItem top="Инклюзивность">
              <Checkbox checked={blind} onChange={handleBlindChange}>
                Доступность для слабовидящих граждан
              </Checkbox>
              <Checkbox checked={wheelchair} onChange={handleWheelchairChange}>
                Доступность для маломобильных граждан
              </Checkbox>
            </FormItem>




          </Group>
        </ModalPage> :
            <ModalPage
                id="filters"
                onClose={modalBack}
                header={
                  <ModalPageHeader


                  >
                    Фильтры банкоматов
                  </ModalPageHeader>
                }
            >
              <Group>





                <FormItem >
                  <Checkbox>Работает сейчас</Checkbox>
                  <Checkbox>NFC для банковских карт</Checkbox>
                  <Checkbox>Снятие наличных с помощью QR кода</Checkbox>
                  <Checkbox>Платежи по QR-коду </Checkbox>
                </FormItem>
                <Separator />
                <FormItem top="Инклюзивность
           ">
                  <Checkbox>Доступность для маломобильных граждан</Checkbox>
                  <Checkbox>Доступность для слабовидящих граждан</Checkbox>
                </FormItem>




              </Group>
            </ModalPage>}
      </ModalRoot>
  );


  return (
      <>
        {/*<JsonFetcher/>*/}
        <AppRoot>
          <SplitLayout modal={modal} header={<PanelHeader separator={false}/>}>
            <SplitCol autoSpaced>
              <View activePanel="main">
                <Panel id="main">
                  <PanelHeader>ВТБ</PanelHeader>

                  <Div>
                  <SegmentedControl
                      name="isOffice"
                      defaultValue= {false}
                      onChange={(value) => changeSelectedType(value)}
                      options={[
                        {
                          label: 'Отделения',
                          value: true,
                        },
                        {
                          label: 'Банкоматы',
                          value: false,
                        },

                      ]}
                  />
                  </Div>
                  <Div>
                     <Button before = {<Icon24Filter/>} onClick={() => changeActiveModal("filters")} appearance={"accent"} stretched = {true}  size ={"l"}>Фильтры</Button>
                  </Div>
                  <YandexMap data={data} coordinates = {coord} me = {MyC}/>
                  <Div style ={{position:"absolute", margin:"5px", bottom:"40px"}}>
                    <Button  appearance={"accent-invariable"} stretched = {true} onClick={() => changeActiveModal("list")} size ={"l"}>Мне повезет</Button>
                  </Div>

                  {/*<CellButton className = {"component"} onClick={() => changeActiveModal("filters")}>*/}
                  {/*  Открыть модальную страницу*/}
                  {/*</CellButton>*/}

                </Panel>
              </View>
            </SplitCol>
          </SplitLayout>
        </AppRoot>
      </>
  );
};

export default App
