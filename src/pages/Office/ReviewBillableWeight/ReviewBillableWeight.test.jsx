import React from 'react';
import { render } from '@testing-library/react';

import ReviewBillableWeight from './ReviewBillableWeight';

describe('ReviewBillableWeight', () => {
  it('renders the component', () => {
    const { getByText } = render(<ReviewBillableWeight />);

    expect(getByText('Review Billable Weight page')).toBeInTheDocument();
  });
});
