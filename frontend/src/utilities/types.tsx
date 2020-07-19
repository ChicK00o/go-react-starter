import {ChangeEvent} from "react";

type InputEvent = ChangeEvent<HTMLInputElement>
export type InputChangeHandler = (e: InputEvent) => void

type SelectEvent = ChangeEvent<HTMLSelectElement>
export type SelectChangeHandler = (e: SelectEvent) => void
