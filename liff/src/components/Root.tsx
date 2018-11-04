import * as React from 'react'
import { Form } from './Form'

export interface RootProps {}
export interface RootState {}

export class Root extends React.Component<RootProps, RootState> {
  constructor(props: RootProps) {
    super(props)
  }

  render() {
    return (
      <div>
        <div>Hello!!</div>
        <Form />
      </div>
    )
  }
}
