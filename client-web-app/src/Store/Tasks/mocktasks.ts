import { Task } from './tasksSlice';

const mocktasks: Task[] = [
  {
    id: '213',
    date: new Date(2022, 5, 1),
    title: 'Залупа ебучая',
    description: 'Сюда я ходил 5 лет почти каждый день',
    time: '12:00',
    temporary: true,
    latitude: 52.262579997411414,
    longitude: 104.26150195437111,
    completed: false,
  },
  {
    id: '3242',
    date: new Date(2022, 5, 6),
    title: 'Хата Зайца',
    description: 'Говорят, тут обитают татары',
    time: '12:00',
    temporary: true,
    latitude: 52.251619911681516,
    longitude: 104.2597259638163,
    completed: false,
  },
  {
    id: '32422',
    date: new Date(2022, 5, 6),
    title: 'dwads Зайца',
    description: 'Говорят, тут обитают татары',
    time: '12:00',
    temporary: true,
    latitude: 52.251619911681516,
    longitude: 104.4597259638163,
    completed: false,
  },
];

export default mocktasks;
