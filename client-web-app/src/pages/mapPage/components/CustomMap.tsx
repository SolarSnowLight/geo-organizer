import React, { useState } from 'react';
import Map, {
  MapboxEvent,
  MapLayerMouseEvent,
  Marker,
  NavigationControl,
  Popup,
} from 'react-map-gl';
import { useActions } from '../../../Store/hooks/useActions';
import '../map.module.sass';
import useTypedSelector from '../../../Store/hooks/useTypedSelector';
import { BarEnum } from '../../../Store/App/appSlice';
import { GeocodeType } from '../../../types/GeocodeTypes';
import Pin from './Pin';

export interface MarkerType {
  longitude: number;
  latitude: number;
}

function CustomMap() {
  const { bar } = useTypedSelector((state) => state.app);
  const { data } = useTypedSelector((state) => state.address);
  const tasks = useTypedSelector((state) => state.tasks.data);
  const { fetchAddress, deleteMarker, deleteAddress } = useActions();
  const [popup, setPopup] = useState<GeocodeType>();
  const handleMapClick = (e: MapLayerMouseEvent) => {
    switch (bar) {
      case BarEnum.NEW_TASK: {
        const longitude = e.lngLat.lng;
        const latitude = e.lngLat.lat;
        fetchAddress({
          longitude,
          latitude,
        });
        break;
      }
      case BarEnum.TASK_LIST: {
        break;
      }
      default: {
        break;
      }
    }
  };
  const handleMarkerClick = (e: MapboxEvent<MouseEvent>) => {
    e.originalEvent.stopPropagation();
    setPopup(data);
    e.target.remove();
    deleteMarker();
    deleteAddress();
  };
  return (
    <Map
      initialViewState={{
        longitude: 104.29739289456052,
        latitude: 52.27600080292447,
        zoom: 13,
      }}
      style={{
        width: '100%',
        height: '100vh',
      }}
      mapboxAccessToken={process.env.REACT_APP_MAPBOX_TOKEN}
      mapStyle="mapbox://styles/mapbox/streets-v11"
      onClick={handleMapClick}
      doubleClickZoom={false}
      id="map"
      attributionControl={false}
    >
      <NavigationControl
        style={{
          display: 'flex',
          alignItems: 'center',
          width: '60px',
          height: '25px',
        }}
        position="bottom-right"
        showCompass={false}
      />
      {data && (
        <Marker
          onClick={handleMarkerClick}
          key={data.features[0].center[0] + data.features[0].center[1]}
          // longitude={data.features[0].geometry.coordinates[0]}
          // latitude={data.features[0].geometry.coordinates[1]}
          longitude={data.query[0]}
          latitude={data.query[1]}
          clickTolerance={20}
        />
      )}
      {tasks
        && tasks.map((tasks) => (
          <Marker
            key={tasks.latitude + tasks.longitude}
            longitude={tasks.longitude}
            latitude={tasks.latitude}
          >
            <Pin />
          </Marker>
        ))}
      {popup && data && (
        <Popup
          longitude={data.features[0].geometry.coordinates[0]}
          latitude={data.features[0].geometry.coordinates[1]}
          anchor="top"
        >
          <div>
            PIDORAS:
            {data.features[0].address}
          </div>
        </Popup>
      )}
    </Map>
  );
}

export default CustomMap;
