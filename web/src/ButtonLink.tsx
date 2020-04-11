import React from "react"
import styled from "styled-components"
import { Button } from "./Button"
import { Font, FontSize, Color, SizeUnit } from "./style-helpers"

let ButtonLinkRoot = styled.a`
  ${Button}
`

type ButtonLinkProps = {
  href: string
  label: string
  target?: string
  rel?: string
}

function ButtonLink(props: ButtonLinkProps) {
  return (
    <ButtonLinkRoot href={props.href} {...props}>
      {props.label}
    </ButtonLinkRoot>
  )
}

export default ButtonLink
