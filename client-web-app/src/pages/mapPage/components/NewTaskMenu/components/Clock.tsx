import React, { ChangeEvent } from 'react';
import TimeField from 'react-simple-timefield';
import useActions from '../../../../../Store/hooks/useActions';

function Clock() {
  const { setTime } = useActions();
  const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
    setTime(e.target.value);
  };
  return (
    <TimeField
      style={{
        width: '40px',
        padding: '0',
        height: '20px',
        borderRadius: '0',
        background: 'transparent',
        marginBottom: '0',
      }}
      onChange={handleChange}
    />
  );
}

export default Clock;
