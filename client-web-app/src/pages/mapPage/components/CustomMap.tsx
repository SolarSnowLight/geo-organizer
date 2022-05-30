import React, { useState } from 'react';
import Map, {
  MapLayerMouseEvent,
  Marker,
  NavigationControl,
} from 'react-map-gl';
import { useActions } from '../../../Store/hooks/useActions';
import '../map.module.sass';

const token = 'pk.eyJ1IjoiaW5ub2tlbnRpeTI1MTciLCJhIjoiY2wwdHRicHd2MHAxZjNibm1odTdwNXk1cCJ9.aibsBxys2tKJkN25qkCAKg';

export interface MarkerType {
  longitude: number;
  latitude: number;
}

function CustomMap() {
  const { fetchAddress, deleteMarker, deleteAddress } = useActions();
  const [markers, setMarkers] = useState<MarkerType[]>([]);
  const onClick = (e: MapLayerMouseEvent) => {
    const longitude = e.lngLat.lng;
    const latitude = e.lngLat.lat;
    fetchAddress({
      longitude,
      latitude,
    });
    setMarkers((markers) => [
      ...markers,
      {
        longitude,
        latitude,
      },
    ]);
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
      mapboxAccessToken={token}
      mapStyle="mapbox://styles/mapbox/streets-v11"
      onDblClick={onClick}
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
      {markers
        && markers.map((marker) => (
          <Marker
            onClick={(e) => {
              e.target.remove();
              deleteMarker();
              deleteAddress();
            }}
            key={marker.longitude}
            longitude={marker.longitude}
            latitude={marker.latitude}
            clickTolerance={20}
          />
        ))}
    </Map>
  );
}

export default CustomMap;
