import React from 'react';
import s from './TaskList.module.sass';
import { Task } from '../../../../Store/Tasks/tasksSlice';

interface ListItem {
  id: string;
  completed: boolean;
  temporary: boolean;
  description: string;
  date: Date;
  weekDay: string[];
  time: string;
}

function ListItem({
  longitude,
  latitude,
  id,
  completed,
  temporary,
  title,
  description,
  date,
  weekDay,
  time,
}: ListItem) {
  return <div className={s.taskitem__wrapper}>{title}</div>;
}

export default ListItem;
