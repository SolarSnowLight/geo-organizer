import React from 'react';
import {
  styled,
  ToggleButton,
  ToggleButtonGroup,
  ToggleButtonProps,
} from '@mui/material';
import useTypedSelector from '../../../../../Store/hooks/useTypedSelector';
import useActions from '../../../../../Store/hooks/useActions';

const DAYS = [
  {
    key: 'Понедельник',
    label: 'ПН',
  },
  {
    key: 'Вторник',
    label: 'ВТ',
  },
  {
    key: 'Среда',
    label: 'СР',
  },
  {
    key: 'Четверг',
    label: 'ЧТ',
  },
  {
    key: 'Пятница',
    label: 'ПТ',
  },
  {
    key: 'Суббота',
    label: 'СБ',
  },
  {
    key: 'Воскресенье',
    label: 'ВС',
  },
];

function WeekDayPicker() {
  const StyledToggle = styled(ToggleButton)<ToggleButtonProps>(() => ({
    '&.Mui-selected': {
      background: '#F3C247',
    },
    '&:hover': {
      background: '#F3C247',
    },
    '&.Mui-selected:hover': {
      background: '#F3C247',
    },
  }));
  const { weekDay } = useTypedSelector((state) => state.newTask);
  const { setWeekDay } = useActions();
  // const [days, setDays] = useState([]);
  return (
    <ToggleButtonGroup
      value={weekDay}
      onChange={(event, value) => setWeekDay(value)}
      style={{
        width: '100%',
        justifyContent: 'space-between',
        marginBottom: '24px',
      }}
    >
      {DAYS.map((day, index) => (
        <StyledToggle
          value={index}
          key={day.key}
          style={{
            border: 'none',
            fontSize: '14px',
            fontWeight: '400',
            borderRadius: '50%',
            width: '45px',
            height: '45px',
            color: '#000000',
          }}
        >
          {day.label}
        </StyledToggle>
      ))}
    </ToggleButtonGroup>
  );
}

export default WeekDayPicker;
