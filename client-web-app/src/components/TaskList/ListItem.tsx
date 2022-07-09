import React from 'react';
import { Checkbox } from '@mui/material';
import s from './TaskList.module.sass';

/* interface ListItem {
  id: string;
  completed: boolean;
  temporary: boolean;
  date: Date;
  weekDay: string[];
} */

interface ListItemType {
  title: string;
  description: string;
  address: string;
  time: string;
}

function ListItem({
  title, description, address, time,
}:ListItemType) {
  return (
    <div className={s.taskitem__wrapper}>
      <Checkbox
        sx={{
          '&.MuiButtonBase-root.MuiCheckbox-root': {
            padding: '0px',
          },
        }}
      />
      <div className={s.taskitem__infowrapper}>
        <span className={s.taskitem__title}>{title}</span>
        <span className={s.taskitem__description}>{description}</span>
        <span className={s.taskitem__address}>{address}</span>
      </div>
      <span className={s.taskitem__time}>{time}</span>
    </div>
  );
}

export default ListItem;
