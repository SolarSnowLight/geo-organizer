import React, { useEffect } from 'react';
import 'react-day-picker/dist/style.css';
import ReactSwitch from 'react-switch';
import logo from '../../../../assets/logo.svg';
import { useTypedSelector } from '../../../../Store/hooks/useTypedSelector';
import s from './NewTaskMenu.module.sass';
import { getAddress } from './utils';
import Clock from './components/Clock';
import WeekDayPicker from './components/WeekDayPicker';
import Calendar from './components/Calendar';
import useActions from '../../../../Store/hooks/useActions';
import { BarEnum } from '../../../../Store/App/appSlice';

function NewTaskMenu() {
  const { data } = useTypedSelector((state) => state.address);
  const state = useTypedSelector((state) => state.newTask);
  const { temporary, title, description } = useTypedSelector(
    (state) => state.newTask,
  );
  const address: string = getAddress(data);
  const {
    setTitle,
    setDescription,
    setTemporary,
    resetState,
    setBar,
    addTask,
    deleteAddress,
    deleteMarker,
  } = useActions();
  useEffect(() => {
    resetState();
  }, []);
  const handleTitleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setTitle(e.target.value);
  };
  const handleDescChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setDescription(e.target.value);
  };
  const handleCancelClick = () => {
    setBar(BarEnum.TASK_LIST);
    deleteMarker();
    deleteAddress();
  };
  const handleCreateClick = () => {
    if (data) {
      addTask({
        longitude: data.query[0],
        latitude: data.query[1],
        title: state.title,
        description: state.description,
        date: state.date,
        weekDay: state.weekDay,
        temporary: state.temporary,
        id: new Date().toString(),
        time: state.time,
      });
      setBar(BarEnum.TASK_LIST);
      deleteMarker();
      deleteAddress();
    }
  };
  return (
    <div className="taskbar newTaskMenu">
      <img src={logo} alt="Logo" />
      <h1>Создание метки</h1>
      <h2>{address}</h2>
      <input
        value={title}
        placeholder="Название"
        onChange={handleTitleChange}
      />
      <input
        value={description}
        placeholder="Описание"
        onChange={handleDescChange}
        className={s.newTaskMenu__desc}
      />
      <div className={s.newTaskMenu__switchWrapper}>
        <p>
          Временная
          {' '}
          <span>(удалится после выполнения)</span>
        </p>
        <ReactSwitch
          checked={temporary}
          onChange={setTemporary}
          width={22}
          height={14}
          handleDiameter={6}
          uncheckedIcon={false}
          checkedIcon={false}
        />
      </div>
      {temporary ? <Calendar /> : <WeekDayPicker />}
      <div className={s.newTaskMenu__timePickWrapper}>
        <span>Время уведомления</span>
        <Clock />
      </div>
      <div className={s.newTaskMenu__buttonWrapper}>
        <button type="button" onClick={handleCancelClick}>
          Отмена
        </button>
        <button
          type="button"
          className={s.newTaskMenu__buttonWrapper_create}
          onClick={handleCreateClick}
        >
          Создать
        </button>
      </div>
    </div>
  );
}

export default NewTaskMenu;
