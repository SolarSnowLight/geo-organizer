import React from 'react';
import useTypedSelector from '../../Store/hooks/useTypedSelector';
import ListItem from './ListItem';

interface DateItemInterface {
  date: string;
}

export const getDate = (date: Date) => {
  const dayNumber: number = date.getDate();
  const months = [
    'Января',
    'Февраля',
    'Марта',
    'Апреля',
    'Мая',
    'Июня',
    'Июля',
    'Августа',
    'Сентября',
    'Октября',
    'Ноября',
    'Декабря',
  ];
  const month: string = months[date.getMonth()];
  return `${dayNumber} ${month}`;
};

function DateItem({ date: dateStr }: DateItemInterface): JSX.Element {
  const tasks = useTypedSelector((state) => state.tasks);
  const date = getDate(new Date(dateStr));
  // console.log(dateStr);
  return (
    <div>
      {date}
      {tasks.data.map((task) => (
        (task.temporary && task.date?.toString() === dateStr) ? (
          <ListItem
            title={task.title}
            key={task.title}
            description={task.description}
            address={task.address}
            time={task.time}
          />
        ) : null
      ))}
    </div>
  );
}

export default DateItem;
