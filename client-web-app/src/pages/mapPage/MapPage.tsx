import React, { useState } from 'react';
import CustomMap from './components/CustomMap';
import NewTaskMenu from './components/NewTaskMenu/NewTaskMenu';
import s from './map.module.sass';
import TaskList from './components/TaskList';

function MapPage() {
  const [isNewTaskBar, setNewTaskBar] = useState<boolean>(false);
  const handleCreateTask = () => {
    setNewTaskBar(!isNewTaskBar);
  };
  return (
    <div className={s.mapPageWrapper}>
      {isNewTaskBar ? <NewTaskMenu /> : <TaskList />}
      <CustomMap />
      <button
        onClick={handleCreateTask}
        className={s.mapPageWrapper__newTaskButton}
        type="button"
      >
        Добавить задание
      </button>
    </div>
  );
}

export default MapPage;
