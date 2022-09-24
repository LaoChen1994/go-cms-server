import useBreadcrumb from "Hooks/useBreadcrumb";

export default function () {
  const breadCrumb = useBreadcrumb()
  return (
    <div>
      {breadCrumb}
      <div>Content</div>
    </div>
  )
}
