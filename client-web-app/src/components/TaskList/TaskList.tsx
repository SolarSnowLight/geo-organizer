import React from 'react';
import {
  Tab, TabList, TabPanel, TabPanels, Tabs,
} from '@chakra-ui/tabs';
import logo from '../../assets/logo.svg';
import searchIcon from '../../assets/search-icon.svg';
import useTypedSelector from '../../Store/hooks/useTypedSelector';
import ListItem from './ListItem';
import DateItem from './DateItem';

const tabs = [{ text: 'Текущие' }, { text: 'Постоянные' }, { text: 'Выполненные' }];

function TaskList() {
  const tasks = useTypedSelector((state) => state.tasks);
  const dates = new Set<string>();
  tasks.data.map((task) => (
    task.date && dates.add(task.date.toString())
  ));
  /* if (dates.size !== 0) {
    dates.forEach((value) => console.log(new Date(value)));
  } */
  const daysComponents: JSX.Element[] = [];
  dates.forEach((value) => (daysComponents.push(<DateItem key={value} date={value} />)));
  return (
    <div className="taskbar">
      <img src={logo} alt="Logo" />
      <h1>Список заданий</h1>
      <div className="search-input">
        <button type="button">
          <img alt="search" src={searchIcon} />
        </button>
        <input placeholder="Поиск" />
      </div>
      <Tabs>
        <TabList className="tabs-switch">
          {tabs.map(((value) => (
            <Tab
              className="tabs-switch__item"
              _selected={{ background: '#fbfbfb' }}
              key={value.text}
            >
              {value.text}
            </Tab>
          )))}
        </TabList>
        <TabPanels>
          <TabPanel>
            {daysComponents.length !== 0 && daysComponents.map((day) => (day))}
            {/* Список текущих заданий
            {tasks.data.map((task) => (
              task.temporary ? (
                <ListItem
                  title={task.title}
                  key={task.title}
                  description={task.description}
                  address={task.address}
                  time={task.time}
                />
              ) : null
            ))} */}
            {/* {
              dates.forEach((value) => (
                <>
                  {}
                  {tasks.data.map((task) => (
                    (task.temporary && task.date === new Date(value)) ? (
                      <ListItem
                        title={task.title}
                        key={task.title}
                        description={task.description}
                        address={task.address}
                        time={task.time}
                      />
                    ) : null
                  ))}
                </>
              ))
            } */}
          </TabPanel>
          <TabPanel>
            Список постоянных заданий
            {tasks.data.map((task) => (
              task.temporary ? null : (
                <ListItem
                  title={task.title}
                  key={task.title}
                  description={task.description}
                  address={task.address}
                  time={task.time}
                />
              )
            ))}
          </TabPanel>
          <TabPanel>Список выполненных заданий</TabPanel>
        </TabPanels>
      </Tabs>
    </div>
  );
}

export default TaskList;
