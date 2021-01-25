/* eslint-disable react/jsx-props-no-spreading */
import React from 'react';
import { mount } from 'enzyme';

import Contact from '.';

const defaultProps = {
  header: '',
  dutyStationName: '',
  officeType: '',
  telephone: '',
};
function mountFooter(props = defaultProps) {
  return mount(<Contact {...props} />);
}
describe('Contact component', () => {
  it('renders footer with given required props', () => {
    const header = 'Contact Info';
    const dutyStationName = 'Headquarters';
    const officeType = 'Homebase';
    const telephone = '(777) 777-7777';
    const props = {
      header,
      dutyStationName,
      officeType,
      telephone,
      moveSubmitted: false,
    };
    const wrapper = mountFooter(props);
    expect(wrapper.find('h6').text()).toBe(header);
    expect(wrapper.find('strong').text()).toBe(dutyStationName);
    expect(wrapper.find('span').length).toBe(2);
    expect(wrapper.find('span').at(0).text()).toBe(officeType);
    expect(wrapper.find('span').at(1).text()).toBe(telephone);
    expect(wrapper.find('[data-testid="move-submitted-instructions"]').exists()).toBe(false);
  });

  it('renders additional copy if the move is submitted', () => {
    const header = 'Contact Info';
    const dutyStationName = 'Headquarters';
    const officeType = 'Homebase';
    const telephone = '(777) 777-7777';
    const props = {
      header,
      dutyStationName,
      officeType,
      telephone,
      moveSubmitted: true,
    };
    const wrapper = mountFooter(props);
    expect(wrapper.find('[data-testid="move-submitted-instructions"]').text()).toBe(
      'Talk to your move counselor or directly with your movers if you have questions during your move.',
    );
  });
});
