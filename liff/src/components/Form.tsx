import * as React from 'react'
import * as liff from '../liff/liff'

export interface FormProps {}
export interface FormState {
  textValue: string
}

interface MessageInputEvent extends React.FormEvent<HTMLInputElement> {
  target: HTMLInputElement
}

export class Form extends React.Component<FormProps, FormState> {
  constructor(props: FormProps) {
    super(props)

    this.state = {
      textValue: ''
    }
  }

  changeText = (e: MessageInputEvent) => {
    const state: FormState = {
      textValue: e.target.value
    }
    this.setState(state)
  }

  render = () => {
    return (
      <div>
        <input
          type="text"
          value={this.state.textValue}
          onChange={this.changeText}
        />
        <button
          onClick={() => {
            liff.sendMessage(this.state.textValue)
          }}
        >
          button
        </button>
      </div>
    )
  }
}
