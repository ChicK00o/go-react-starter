import {ChangeEvent} from "react";

type InputEvent = ChangeEvent<HTMLInputElement>
export type ChangeHandler = (e: InputEvent) => void
