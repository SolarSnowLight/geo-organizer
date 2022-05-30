import React from 'react';
import {
  Tab, TabList, TabPanel, TabPanels, Tabs,
} from '@chakra-ui/tabs';
import logo from '../../../assets/logo.svg';
import searchIcon from '../../../assets/search-icon.svg';

const tabs = [{ text: 'Текущие' }, { text: 'Постоянные' }, { text: 'Выполненные' }];

function TaskList() {
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
          <TabPanel>Список текущих заданий</TabPanel>
          <TabPanel>Список постоянных заданий</TabPanel>
          <TabPanel>Список выполненных заданий</TabPanel>
        </TabPanels>
      </Tabs>
    </div>
  );
}

export default TaskList;
